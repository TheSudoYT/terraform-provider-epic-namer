# terraform-provider-epic

The most EPIC provider for generating the most EPIC resource names, descriptions, and tags.

## Example Usage

Name an S3 bucket after your favorite Lord of the Rings character!

```hcl
terraform {
  required_providers {
    epic = {
      version = ">=0.1.0"
      source  = "TheSudoYT/epic"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">=5.47.0"
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

resource "epic_random_quote" "lotr_quote" {
  media_type = "movie"
  title      = "lord of the rings"
}


resource "aws_s3_bucket" "epic" {
  bucket = epic_random_name.movie_name.name

  tags = {
    Name        = epic_random_name.movie_name.name
    Description = epic_random_quote.lotr_quote.quote
  }
}
```

## Media Types and Titles

Media types are mediums that a title is nested within. Titles are the titles of movies, video games, etc.

The following `media_type`s and their associated `title`s are listed below:

```
Media Type: anime
Titles:
  - one_piece
  - spy_x_family

Media Type: movie
Titles:
  - jurassic_park
  - lord_of_the_rings
  - star_wars

Media Type: tv_series
Titles:
  - breaking_bad
  - game_of_thrones

Media Type: video_game
Titles:
  - final_fantasy_vii
  - kingdom_hearts_1
  - the_witcher
```
