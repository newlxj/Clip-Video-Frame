/*
*
create by newlxj
2024-07-11 21:44
*
*/
package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var videoDir = "./"
var outTempDir = "outtemp" //临时目录
var newOutDir = "newout"

var videoType = "gif" //输出文件类型支持mp4，gif等格式
var frameNum = "30"   //拆帧数 视频拆成多少帧，会取视频总时长平均值
var videoSpeed = "2"  //gif速度(1,2,3,4,5)
var videoSize = "256" //高宽缩放大小（比如 256）

// 定义格式名称数组
var extensions = []string{
	".avi",
	".mp4",
	".mov",
	".mpg",
	".mpeg",
	".m4v",
	".mkv",
	".webm",
	".flv",
	".3gp",
	".3g2",
	".asf",
	".wmv",
	".wma",
	".ogg",
	".ogv",
	".ts",
	".m2ts",
	".mts",
	".mxf",
}

func getPath(fileName string) string {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		// return ""
	}
	// Get the directory of the current executable
	exeDir := filepath.Dir(exePath)
	// Construct the path to ffmpeg

	os := runtime.GOOS
	switch os {
	case "windows":
		fileName = "fg/win/bin/" + fileName
	case "darwin":
		fileName = "fg/linux/bin/" + fileName
	case "linux":
		fileName = "fg/linux/bin/" + fileName
	default:
		log.Fatal("Current OS is unknown")
	}

	filex := filepath.Join(exeDir, fileName)
	return filex
}

// 获取视频总时长（以秒为单位）
func getVideoDuration(videoPath string) (float64, error) {
	cmd := exec.Command(getPath("ffprobe.exe"), "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", videoPath)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(string(output))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}

// 提取指定时间的帧并保存为 JPG
func extractFrame(videoPath, outputImagePath string, startTime float64, wg *sync.WaitGroup) {

	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("提取帧 %s 时发生错误: %v", outputImagePath, r)
		}
	}()
	startTimeStr := fmt.Sprintf("%.2f", startTime)

	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, getPath("ffmpeg"), "-ss", startTimeStr, "-i", videoPath, "-vf", "scale="+videoSize+":-1", "-vframes", "1", "-q:v", "0", outputImagePath)
	err := cmd.Run()
	if err != nil {
		log.Printf("提取帧 %s 失败: %v", outputImagePath, err)
	} else {
		fmt.Printf("提取帧: %s (开始时间: %f 秒)\n", outputImagePath, startTime)
	}

	// 检查是否由于上下文超时而失败
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("提取帧 %s 失败: 超时", outputImagePath)
	}
}

// 合并 JPG 图片为 GIF，设置每帧播放速度
func createGifFromImages(imagePattern, outputFilePath string, delayStr string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, getPath("ffmpeg"), "-y", "-framerate", videoSpeed+"/"+delayStr, "-i", imagePattern, "-vf", "scale="+videoSize+":-1,setpts=PTS/"+delayStr, outputFilePath)
	cmd.Run()
	// 检查是否由于上下文超时而失败
	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("提取帧 %s 失败: 超时", imagePattern)
	}
	return nil
}

// 删除临时 JPG 文件
func deleteTempImages(imagePattern string) {
	files, err := filepath.Glob(imagePattern)
	if err != nil {
		log.Printf("查找临时图片文件出错: %v", err)
		return
	}
	// log.Printf("找到的文件: %v", files)

	for _, file := range files {
		os.Remove(file)
		// if err != nil {
		// 	log.Printf("删除临时图片文件 %s 失败: %v", file, err)
		// } else {
		// 	log.Printf("成功删除临时图片文件 %s", file)
		// }
	}
}

// 生成唯一 GIF 文件名
func generateUniqueGifName(basePath, videoPath string) string {
	videoPath = strings.ReplaceAll(videoPath, string(filepath.Separator), "_")
	videoPath = strings.ReplaceAll(videoPath, ".", "_")
	videoPath = strings.ReplaceAll(videoPath, ":", "_")
	videoPath = strings.ReplaceAll(videoPath, "&", "_")

	os := runtime.GOOS
	switch os {
	case "windows":
		if len(videoPath) > 800 {
			videoPath = videoPath[:800]
		}
	case "darwin":
		if len(videoPath) > 80 {
			videoPath = videoPath[:80]
		}
	case "linux":
		if len(videoPath) > 80 {
			videoPath = videoPath[:80]
		}
	default:
		fmt.Println("Current OS is unknown")
	}

	outputFilePath := filepath.Join(basePath, "ClipVideoFrame_"+videoPath+"."+videoType)
	// 如果文件已存在，添加随机数后缀
	randNum, _ := rand.Int(rand.Reader, big.NewInt(10000))
	base := strings.TrimSuffix(filepath.Base(outputFilePath), "."+videoType)
	dir := filepath.Dir(outputFilePath)
	outputFilePath = filepath.Join(dir, fmt.Sprintf("%s_%d."+videoType, base, randNum))

	// for {
	// 	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
	// 		break
	// 	}

	// }

	return outputFilePath
}

func goWork() {

	infoFilePath := "info.txt"
	// 创建输出目录
	if err := os.MkdirAll(outTempDir, os.ModePerm); err != nil {
		log.Fatalf("创建目录 %s 失败: %v", outTempDir, err)
	}
	if err := os.MkdirAll(newOutDir, os.ModePerm); err != nil {
		log.Fatalf("创建目录 %s 失败: %v", newOutDir, err)
	}

	// 读取已存在的日志记录
	existingRecords := make(map[string]bool)

	if _, err := os.Stat(infoFilePath); err == nil {
		infoFile, err := os.Open(infoFilePath)
		if err != nil {
			log.Fatalf("读取 info.txt 文件失败: %v", err)
		}
		defer infoFile.Close()

		scanner := bufio.NewScanner(infoFile)
		for scanner.Scan() {
			record := scanner.Text()
			parts := strings.Split(record, ";")
			if len(parts) == 2 {
				existingRecords[parts[0]] = true
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("读取 info.txt 文件出错: %v", err)
		}
	}

	infoFile, err := os.OpenFile(infoFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("创建/打开 info.txt 文件失败: %v", err)
	}
	defer infoFile.Close()

	err = filepath.Walk(videoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if strings.Contains(strings.Join(extensions, "|"), strings.ToLower(ext)) {
				if existingRecords[path] {
					fmt.Printf("已处理文件: %s，跳过...\n", path)
					return nil
				}

				videoPath := path
				duration, err := getVideoDuration(videoPath)
				if err != nil {
					log.Printf("获取视频时长出错: %v", err)
					return nil // 继续处理下一个文件
				}

				fmt.Printf("视频总时长: %f 秒\n", duration)

				// 计算每个帧的开始时间
				frameCount := 30
				interval := duration / float64(frameCount+1)

				var wg sync.WaitGroup
				wg.Add(frameCount)

				// 获取当前时间的字符串表示
				currentTime := time.Now().Format("20060102150405")
				// 生成10位随机数
				randomNumber, _ := rand.Int(rand.Reader, big.NewInt(10000000))

				// 组合当前时间和10位随机数
				videoName := fmt.Sprintf("%s%010d", currentTime, randomNumber)

				// // 检查字符串的长度，避免越界
				// if len(videoName) > 80 {
				// 	videoName = videoName[:80]
				// }

				// 并行提取帧并保存为 JPG
				for i := 0; i < frameCount; i++ {
					startTime := interval * float64(i+1)
					outputImagePath := fmt.Sprintf("%s/%s_frame_%03d.jpg", outTempDir, videoName, i+1)
					// wg.Add(1)
					go extractFrame(videoPath, outputImagePath, startTime, &wg)
				}

				wg.Wait()
				// 合并 JPG 图片为 GIF
				imagePattern := fmt.Sprintf("%s/%s_frame_%%03d.jpg", outTempDir, videoName)
				outputFilePath := generateUniqueGifName(newOutDir, videoPath)

				err = createGifFromImages(imagePattern, outputFilePath, "0.002")
				if err != nil {
					log.Printf("合并 JPG出错: %v", err)
					deleteTempImages(fmt.Sprintf("%s/%s_frame_*.jpg", outTempDir, videoName))
					return nil // 继续处理下一个文件
				}
				fmt.Printf("生成文件: %s\n", outputFilePath)

				// 删除临时 JPG 文件
				deleteTempImages(fmt.Sprintf("%s/%s_frame_*.jpg", outTempDir, videoName))

				// 记录信息到 info.txt 文件
				_, err = fmt.Fprintf(infoFile, "%s;%s\n", videoPath, outputFilePath)
				if err != nil {
					log.Printf("写入 info.txt 文件失败: %v", err)
					return nil // 继续处理下一个文件
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("扫描目录出错:可能输入的路径不存在", err)
	}
}

func main() {
	fmt.Println("\r\n\r\n\r\n\r\n\r\n\r\n\r\n")
	fmt.Println("视频文件切抽帧剪切工具，可以将视频拆成多帧并合并成mp4、gif输出")
	fmt.Println("可以通过传入参数方式运行:cutvideo ./ ./outvideo 30 2 gif 256")
	fmt.Println("参数:   文件夹位置 输出文件夹位置  拆帧数  视频速度(1,2,3,4,5) 输出文件格式(gif,mp4) 高宽缩放大小（比如 256）")
	fmt.Println("\r\n\r\n\r\n\r\n\r\n\r\n\r\n")
	//传入参数:   文件夹位置 输出文件夹位置  拆帧数  视频速度(1,2,3,4,5) 输出文件格式(gif,mp4) 高宽缩放大小（比如 256）
	args := os.Args
	if len(args) > 6 {
		videoDir = args[1]
		newOutDir = args[2]
		frameNum = args[3]
		videoSpeed = args[4]
		videoType = args[5]
		videoSize = args[6]
	} else {
		fmt.Println("请输入相关参数进行视频抽帧转换：")
		fmt.Println(fmt.Sprintf("文件夹路径(默认:%s): ", videoDir))
		fmt.Scanln(&videoDir)
		fmt.Println(fmt.Sprintf("输出文件夹(默认:%s): ", newOutDir))
		fmt.Scanln(&newOutDir)
		fmt.Println(fmt.Sprintf("拆帧数(默认:%s): ", frameNum))
		fmt.Scanln(&frameNum)
		fmt.Println(fmt.Sprintf("视频每帧播放速度(1,2,3,4,5 默认:%s): ", videoSpeed))
		fmt.Scanln(&videoSpeed)
		fmt.Println(fmt.Sprintf("视频输出类型(mp4,gif,默认:%s): ", videoType))
		fmt.Scanln(&videoType)
		fmt.Println(fmt.Sprintf("高宽缩放大小(默认:%s): ", videoSize))
		fmt.Scanln(&videoSize)
	}

	goWork()
	fmt.Println("任务执行完成请检查输出文件夹", newOutDir)
	time.Sleep(240 * time.Second)
}
