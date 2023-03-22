## Technical Assignment Completion

**This application allows you to upload and store images tied to a specific user. It was built using clean architecture
and SOLID principles for future scalability.**

- Folder "cmd" contains the main entry point of the application.
- Folder "config" contains environment variables necessary for the application to function.
- Folder "internal" contains the "app" folder which has the file that starts all the necessary dependencies for the
  application to function. Migrations are also located here.
    - Folder "controller" contains everything for receiving data from the internet, including middleware for basic
      authorization, handlers, routers, and entities for requests and responses.
    - Folder "database" contains the file for working with the database.
    - All entities are located in the "entity" folder.
    - The services are located in the "service" folder.
- In the "migrations" folder, there are SQL schemas of entities
- The 'pkg' folder contains files with external dependencies for implementing clean architecture. They are replaced with
  interfaces and structures that implement them, allowing for easy substitution at any time

## Test Routes

- **To run the project, utilize the docker-compose file located in the ".docker" folder.**
- **For testing purposes, you can make use of the Postman collection in the ".postman" folder.**

## Endpoints

1. **Register user. A test user is automatically added to the database {"username":"TEST USER", "password": "
   1234567890"}, but if you want to create your own, I have added
   a corresponding endpoint.**
    - **Method: POST**
    - ***Url: http://localhost:8080/api/v1/users/register***
        - **Request structure:**
          ```json
          {
            "username": "MY USER",
            "password": "1234567890"
          }
          ```

2. **Login user**

    - **Method: POST**
    - ***Url: http://localhost:8080/api/v1/users/login***
        - **Request structure:**
          ```json
          {
            "username": "TEST USER",
            "password": "1234567890"
          }
          ```

3. **Upload Image**
    - **Method: POST**
    - ***Url: http://localhost:8080/api/v1/images/upload-picture***
    - ***When uploading, you need to provide an image in a format (for example, JPEG or PNG).***

4. **Get all images**
    - **Method: GET**
    - ***Url: http://localhost:8080/api/v1/images***
    - ***You can open any image by image_url from response.***

створити сервіс у якому

- POST /login {"username": "test", "password": 123456} Response: {"token": "jwt string"} юзер зможе залогінитись(
  припустимо що юзер уже є в бд, і для цього треба буде ствоирит запис в бд вручну) та отримати jwt токен який живе 12
  годин. усі логіни зберігати в табличці users [id, username, password_hash]
- POST /upload-picture - завантажити фото, та зберегти її у файл, а урл зберегти в базу даних в табличку
  images [id, user_id, image_path, image_url]
- GET /images - отримати масив усіх своїх картинок

/upload-picture
/images

має мати middleware який буде перевіряти чи валідний токен той який юзер отримав при авторизації. Токен має передатись в
хедері Authorization: Bearer {jwt token here}
