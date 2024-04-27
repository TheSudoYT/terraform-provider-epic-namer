terraform {
  required_providers {
    epic = {
      version = "0.1.0"
      source  = "localhost/providers/epic-namer"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "5.47.0"
    }
  }
}

provider "epic" {}

provider "aws" {
  region = "us-east-1"
}

resource "epic_random_name" "movie_name" {
  media_type = "movie"
  title      = "lord of the rings"
}

resource "epic_random_name" "tv_series_name" {
  media_type = "tv_series"
  title      = "game of thrones"
}

resource "epic_random_quote" "lotr_quote" {
  media_type = "movie"
  title      = "lord of the rings"
}

resource "epic_random_quote" "got_quote" {
  media_type = "tv_series"
  title      = "game of thrones"
}

resource "aws_s3_bucket" "epic" {
  bucket = epic_random_name.movie_name.name

  tags = {
    Name        = epic_random_name.movie_name.name
    Description = epic_random_quote.lotr_quote.quote
  }
}

output "s3_bucket_name" {
  value = aws_s3_bucket.epic.bucket
}

output "random_movie_name" {
  value = epic_random_name.movie_name.name
}

output "random_tv_series_name" {
  value = epic_random_name.tv_series_name.name
}

output "lotr_quote" {
  value = epic_random_quote.lotr_quote.quote
}

output "got_quote" {
  value = epic_random_quote.got_quote.quote
}