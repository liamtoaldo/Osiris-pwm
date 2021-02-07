[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/Gyro7/Osiris-pwm">
    <img src="https://i.imgur.com/gwqaZhK.png" alt="Logo" heigth="60">
  </a>

  <h3 align="center">Osiris Password Manager</h3>

  <p align="center">
    A simple and lightweight encrypted password manager written in Go.
    <br />
    <br />
    <a href="https://github.com/Gyro7/Osiris-pwm/issues">Report Bug</a> || 
    <a href="https://github.com/Gyro7/Osiris-pwm/pulls">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

-   [About the Project](#about-the-project)
    -   [Built With](#built-with)
-   [Getting Started](#getting-started)
    -   [Prerequisites](#prerequisites)
    -   [Installation](#installation)
-   [Usage](#usage)
-   [Roadmap](#roadmap)
-   [Contributing](#contributing)
-   [License](#license)
-   [Contact](#contact)
-   [Acknowledgements](#acknowledgements)

## About The Project  

<br>
<p align="center">You don't have to remember your passwords anymore with Osiris!
  <br>
  <br>
<img src="https://i.imgur.com/pJpXZ6S.png" alt="example" width="800">
</p>

### Built With

-   [Go](https://golang.org)
-   [Fyne GUI](https://fyne.io/)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

-   linux or windows (not available for mac at the moment)

### Installation

#### Debian and Ubuntu-based
```sh
# go to Releases and download the deb package
# install the deb package
sudo apt install ./Osiris-pwm1.0.deb
```
#### Any other Linux distribution
```sh
# select where to download the program, for example here I download it to the Desktop
cd Desktop
# clone and go into repo
git clone https://github.com/Gyro7/Osiris-pwm.git
cd Osiris-pwm/
# change mod
chmod +x Osiris-pwm
# run
./Osiris-pwm
```
#### Any other Linux distribution (Releases method)
```sh
# go to Releases and download the .zip file
# decompress it
unzip -q Osiris-pwm1.0.zip
# go into the directory
cd Osiris-pwm
# change mod
chmod +x Osiris-pwm
# run
./Osiris-pwm
```
#### Any other Linux distribution (Build from source)
For this method you need to have Go installed, build-essential and a few dependencies (even tho they should be downloaded automatically)  
If you encounter any problems in installing dependencies, follow this guide: https://github.com/go-gl/glfw
```sh
# clone and go into repo
git clone https://github.com/Gyro7/Osiris-pwm.git
cd Osiris-pwm/
# remove the linux executable
rm Osiris-pwm
# build
go build
# run
./Osiris-pwm
```
#### Windows
```sh
Go to Releases and download the setup for Windows (.exe file)
Follow the simple steps
IF THE PROGRAM DOESN'T WORK, RIGHT CLICK ON IT, GO TO PROPERTIES, GO TO COMPATIBILITY AND SELECTE "ALWAYS EXECUTE AS ADMINISTRATOR"
```
## Usage

If you followed the previous steps, you just have to run the program.
To add a new element to the grid click the white add button on the middle left.
To delete an element, edit it so that all its entries are empty, then head to the Edit menu and click "Delete empties"
To change theme go to File -> Settings


<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/Gyro7/Osiris-pwm/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

gyro - [gyro@sach1.tk](mailto:gyro@sach1.tk)
Project Link: [https://github.com/Gyro7/Osiris-pwm](https://github.com/Gyro7/Osiris-pwm)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

-   [Myself for doing everything.](https://github.com/Gyro7)
-   [The Fyne library for making it easier](https://fyne.io/)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/Gyro7/Osiris-pwm.svg?style=flat-square
[contributors-url]: https://github.com/Gyro7/Osiris-pwm/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Gyro7/Osiris-pwm.svg?style=flat-square
[forks-url]: https://github.com/Gyro7/Osiris-pwm/network/members
[stars-shield]: https://img.shields.io/github/stars/Gyro7/Osiris-pwm.svg?style=flat-square
[stars-url]: https://github.com/Gyro7/Osiris-pwm/stargazers
[issues-shield]: https://img.shields.io/github/issues/Gyro7/Osiris-pwm.svg?style=flat-square
[issues-url]: https://github.com/Gyro7/Osiris-pwm/issues
[license-shield]: https://img.shields.io/github/license/Gyro7/Osiris-pwm.svg?style=flat-square
[license-url]: https://github.com/Gyro7/Osiris-pwm/blob/main/LICENSE
