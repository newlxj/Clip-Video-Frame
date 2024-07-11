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
![Description](https://example.com/path/to/your/gif.gif)


512px 30 frame  gif
![Description](https://example.com/path/to/your/gif.gif)

256px 30 frame  mp4




