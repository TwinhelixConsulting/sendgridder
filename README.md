# Sendgridder

Sendgridder is a simple Go application that sends emails using the SendGrid API. This application reads email details from environment variables and sends an email with optional HTML content and CC recipients.

## Prerequisites

- Go 1.23.1
- Docker

## Installation

### Using Go

1. Clone the repository:
  ```sh
  git clone https://github.com/TwinhelixConsulting/sendgridder.git
  cd sendgridder
  ```

2. Build the application:
  ```sh
  go build -o sendgridder
  ```

### Using Docker

1. Build the Docker image:
  ```sh
  docker build -t sendgridder .
  ```

## Usage

### Using Go

1. Set the required environment variables:
  ```sh
  export API_KEY="your_sendgrid_api_key"
  export SENDER="sender@example.com"
  export SUBJECT="Your Subject"
  export RECIPIENT="recipient@example.com"
  ```

2. Optionally, set the following environment variables:
  ```sh
  export HAS_HTML="true" # if you want to include HTML content
  export CC="cc@example.com" # if you want to add a CC recipient
  ```

3. Create the email body files:
  ```sh
  echo "This is the plain text body" > /tmp/body.txt
  echo "<p>This is the HTML body</p>" > /tmp/body.html
  ```

4. Run the application:
  ```sh
  ./sendgridder
  ```

### Using Docker

1. Run the Docker container with the required environment variables:
  ```sh
  docker run -e API_KEY="your_sendgrid_api_key" \
         -e SENDER="sender@example.com" \
         -e SUBJECT="Your Subject" \
         -e RECIPIENT="recipient@example.com" \
         -v /tmp/body.txt:/tmp/body.txt \
         -v /tmp/body.html:/tmp/body.html \
         sendgridder
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Acknowledgements

- [SendGrid Go Library](https://github.com/sendgrid/sendgrid-go)
- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)
