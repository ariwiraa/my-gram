# My Gram

MyGram is a personal backend project that draws inspiration from the functionalities of Instagram. This social media platform is designed to allow users to showcase their photos, receive comments and likes from other users, and follow other individuals.

Key Features:

- Photo Sharing: Users can upload and share their favorite photos with the community.
- Interactivity: Enable commenting and liking functionality for users to engage with each other's content.
- Follow System: Users can follow and be followed by other members, creating a connected and dynamic community.

## Tech Stack

- Go
- Gin
- PostgreSQL
- GORM
- JWT
- Redis
- Cloudinary

## Run Locally

Clone the project

```bash
  git clone https://github.com/ariwiraa/my-gram.git
```

Go to the project directory

```bash
  cd my-gram
```

Install dependencies

```bash
  go get -u ./...
```

Start the server

```bash
  go run main.go
```

## Configuring Environment (.env)

This project utilizes configuration through the .env file. To configure your project, follow these steps:

1. Open the .env.example file:

   ```bash
   cp .env.example .env
   ```

2. Open the newly created .env file using your preferred text editor.

3. Adjust the values of environment variables according to your project's requirements. Some common variables that may need to be configured include:

   ```env
   PORT=8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=myuser
   DB_PASSWORD=mypassword
   DB_NAME=mydatabase
   ```

Ensure to review each variable and modify them based on your system or service configurations.

4. Save the changes to the .env file.

Note: Make sure not to include the .env file in your repository and treat it as a sensitive file. The .env file should be ignored in the .gitignore file to avoid it being included in version control.

## Documentation

Since this has not been uploaded to the server yet, while to run it on local first

## How To Contribute

Feel free to contribute to the project by submitting issues, providing feedback, or creating pull requests. Your input is highly valued in making MyGram a vibrant and user-friendly social media platform.

## Explore The Project

Check out the [Github Repository](github.com/ariwiraa/mygram) to explore the project's source code and documentation.

## Get In Touch

I'm open to discussions and collaborations! Connect with me if you have ideas, suggestions, or would like to contribute to the project. Let's make MyGram even better together!

## Authors

- [Linkedin - Ari Wira Atmaja](linkedin.com/in/ari-wira-atmaja)
