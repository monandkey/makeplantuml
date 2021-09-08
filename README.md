
# MakePlantUML

## Overview

This is a tool for quickly creating UML from PCAP.


## Environment

It is designed to be used in the following environments

| OS         |
| :-         |
| Windows 10 |

## Preparation

### Step. 1

Download the software listed in the table below.

| Name | Note | URL |
| :-   | :-   | :-  |
| makeplantuml | Download the zip file from the release. | [Github Release](https://github.com/monandkey/makeplantuml/releases) |
| java |Download it from the official website. | - |
| wireshark |Download it from the official website. | [Download Wireshark](https://www.wireshark.org/download.html) |
| plantuml | Download it from the official website. |[Last version: plantuml.jar](http://sourceforge.net/projects/plantuml/files/plantuml.jar/download) |


<br>

The software should be arranged as follows.

`Java` and `Wireshark` should follow the installer.

```bash
+--- makeplantuml
|   +--- makeplantuml.exe
|   +--- profile
|   |   +--- hosts
|   +--- ext
|   |   +--- plantuml.jar
|   +--- puml
|   +--- pcap
|   +--- result
```


### Step .2

Create a configuration file to use makeplantuml.

The configuration file will be created in your home directory.

```bash
$ makeplantuml.exe init --java-path [PATH] \
                    --plantuml-path [PATH] \
                    --wireshark-path [PATH] \
                    --feature-timestamp false \
                    --feature-name-resolution false \
```

NOTE 1:

`PATH` should contain the path to the software you downloaded.

NOTE 2:

Please enclose the path in double quotation marks.

> "hoge/hoge/hoge"

<br>

After creating the config file, verify that the path is correct.

```bash
$ makeplantuml.exe init --validation-config
OK
$
```

If you see OK, there is no problem.

Where an ERROR is displayed, please check that the path is correct.


### Step .3

Store the PCAP files that you want to convert to UML under PCAP.

```bash
+--- makeplantuml
|   +--- pcap
|   |   +--- hoge.pcap
|   |   +--- fuga.pcap
```

### Step .4

Specify a PCAP file and convert it to UML.

The `PUML` file will be created under `puml` and the `SVG` under `result`.

```bash
$ makeplantuml.exe -f pcap/hoge.pcap
```

## Usage

You can add a timestamp by specifying `-t` or `--timestamp` .

```bash
$ makeplantuml.exe -f pcap/hoge.pcap --timestamp
```

