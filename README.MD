# ASCII-GEN


# 💻 Tech Stack:
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white) ![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)


## DESCRIPTION
ASCII-GEN is a dynamic web application designed for creating artistic ASCII representations from user-provided text inputs. It offers a variety of banner styles to transform plain text into visually appealing ASCII ART. Users can effortlessly generate, customize, and export their ASCII ART creations in both text and document formats.

## AUTHORS
- [Ali Hasan Jasim](https://github.com/AliHJMM)
- [Habib Mansoor](https://github.com/7abib04)
- [Husain Ali](https://github.com/hujaafar)

## Usage

### How to Run
1. Clone the Repo
2. Navigate to the project directory
3. Run `go run main.go`
4. Open `http://localhost:2004` in your web browser.

### Implementation Details

#### Algorithm

ASCII-GEN is created by using ASCII banner templates stored as text files. Each template provides unique ASCII representations for characters. The server reads these templates, converts characters into ASCII ART, and creates the artwork using the user's text input.

#### HTTP Endpoints

- `GET /`: Serves the main HTML page where users can input text and select a banner.
- `POST /ascii-art`: Processes the input text and banner selection, generates the ASCII ART, and displays the result.
- `POST /export`: Facilitates exporting ASCII ART creations in either plain text (txt) format for user download.


#### HTTP Status Codes

- `200 OK`: The Request was Successful
- `400 Bad Request`: The Request was invalid (e.g., missing text or banner).
- `404 Not Found`: The Requested resources was not found.
- `500 Internal Server Error` : An Error occured on the server.
