resource "aws_instance" "web" {
  ami           = "ami-84175ae8"
  instance_type = "t2.micro"
}
