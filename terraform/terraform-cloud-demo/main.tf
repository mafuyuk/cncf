resource "aws_instance" "web" {
  ami           = "ami-0ff21806645c5e492"
  instance_type = "t2.micro"
}
