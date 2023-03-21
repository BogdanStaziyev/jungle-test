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
