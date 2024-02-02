# نویسه
## نویسه چیست
- نویسه یک پروتایپ از یک شبکه مجازی برای اشتراک گذاری متن است
- چیزی شبیه تویتر
## تکنولوژی های استفاده شده
- زبان go و کتابخانه chi router برای منطق برنامه 
- کتابخانه gorm برای برقراری ارتباط با دیتابیس
- استفاده از سالت و تابع sha256 برای هش کردن اطلاعات و تایید درستی انها هنگام ورود 
- استفاده از jwt  همراه با cookies برای اهراز هویت شخص وارد شده 
- کتابخانه htmx برای برقراری ارتباط بین فرانت اند و بک اند
# نحوه استفاده
```
git clone https://github.com/Torbatti/nevise.git
cd nevise
go run main.go
```
# .env
```shell
  PORT=
  AUTH_SECRET=
  JWT_SECRET=
```
# Models
# Show case

- صفحه ی ثبت نام
![Alt text](/showcase/ss1.png)

- صفحه ی ورود
![Alt text](/showcase/ss2.png)

- صفحه ی اصلی
![Alt text](/showcase/ss3.png)

- ایجاد نویسه جدید
![Alt text](/showcase/ss4.png)

- اسکرول بی نهایت با استفاده از htmx revealed
![Alt text](/showcase/ss5.png)
