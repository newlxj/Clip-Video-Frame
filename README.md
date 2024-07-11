 <div align="center">

# Clip-Video-Frame

**English** | [中文](README_ZH_CN.md)

</div>

Super fast frame-by-frame editing and merging into new videos. By extracting the key frames of the video, a 1-hour video can be converted into a 10-second to 1-minute quick preview video, so that you can quickly understand the content of your video.

You can adjust the number of frames to cut (cut according to the average length of the video), the speed of each frame playback, and the maximum height of the generated clip video (proportional scaling).

Supported video formats:
```
	.avi",
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
```

Usage Command line input:
```
##cutvideo Input folder location Output folder location Frame rate Video speed (1,2,3,4,5) Output file format (gif,mp4) Height and width scaling (e.g. 256)##
cutvideo ./ ./outvideo 30 2 gif 256
```
Or directly execute the input parameters
 
 
Let's take a look at the wonderful editing video below:
 
256px 30 frame  gif

![Description](https://github.com/newlxj/Clip-Video-Frame/blob/main/image/CriminalMindsSeason1720p_Criminal_Minds_S01E05_720p_WEB-DL_Dolby_Digital_5_1_h264_mkv_4346.gif?raw=true)

512px 30 frame  gif

![Description](https://github.com/newlxj/Clip-Video-Frame/blob/main/image/_rr_your_video_mp4_6739.gif?raw=true)

![Description](https://github.com/newlxj/Clip-Video-Frame/blob/main/image/CriminalMindsSeason1720p_Criminal_Minds_S01E01_720p_WEB-DL_Dolby_Digital_5_1_h264_mkv_2224.gif?raw=true)

![Description](https://github.com/newlxj/Clip-Video-Frame/blob/main/image/CriminalMindsSeason1720p_Criminal_Minds_S01E02_720p_WEB-DL_Dolby_Digital_5_1_h264_mkv_4816.gif?raw=true)


256px 30 frame  mp4
 
<video id="video" controls="" preload="none" poster="">
      <source id="mp4" src="https://github.com/newlxj/Clip-Video-Frame/blob/main/image/CriminalMindsSeason1720p_Criminal_Minds_S01E04_720p_WEB-DL_Dolby_Digital_5_1_h264_mkv_8374.mp4?raw=true" type="video/mp4">
</videos>
