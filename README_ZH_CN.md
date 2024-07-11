 <div align="center">

# Clip-Video-Frame

[English](README.md) |**中文** 

</div>

超快速逐帧编辑并合并成新视频。通过提取视频的关键帧，可以将一个小时的视频转换为10秒到1分钟的快速预览视频，让您可以快速了解视频内容。



您可以调整剪切的帧数（根据视频的平均长度剪切）、每帧播放的速度以及生成的剪辑视频的最大高度（按比例缩放）。

支持的视频格式：
```
	.avi",
	.mp4",
	.mov",
	.mpg",
	.mpeg",
	.m4v",
	.mkv",
	.webm",
	.flv",
	.3gp",
	.3g2",
	.asf",
	.wmv",
	.wma",
	.ogg",
	.ogv",
	.ts",
	.m2ts",
	.mts",
	.mxf",
```

使用命令行输入：
```
##cutvideo 输入文件夹位置 输出文件夹位置 帧率 视频速度 (1,2,3,4,5) 输出文件格式 (gif,mp4) 高度和宽度缩放 (例如 256)##
cutvideo ./ ./outvideo 30 2 gif 256
```
或者直接执行输入参数

 
让我们来看看下面精彩的编辑视频：

256px 30 帧  gif
![描述](https://example.com/path/to/your/gif.gif)

512px 30 帧  gif
![描述](https://example.com/path/to/your/gif.gif)

256px 30 帧  mp4