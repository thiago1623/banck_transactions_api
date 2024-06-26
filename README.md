# Challenge Stack
This application is a sample REST API for the challenge of handling transactions using a.csv file.
This REST API send the summary information about all transactions of this file, and save this file in bucket s3 of aws


# Methodologies that were used

This API was built with the MVC architecture pattern,
also the repository pattern, to decouple the API logic from the implementation details of data persistence, leading to more modular, maintainable and testable code.
SOLID principles were used, specifically the SRP, OCP and ISP principles


# Getting started

So that you can quickly run through the challenge, some bootstrap scripts have been created to make things easier.

but first, download the file called settings.ini and add it inside config folder for example it has to look like this: config/settings.ini

file link: https://drive.google.com/drive/folders/1F1kX3gSHTzyg3pBN3qH_HmldLzyk6j-Y?usp=sharing

> change the sender_email, recipient_email and email_password settings to your own settings so you can validate if you received the email.


---

In plain language, all you need to do is run bootstrap with make to build the container.


A detailed step-by-step description is:
```
make build
```
The development server should have started now. You can visit the API by navigating in a browser to: `http://0.0.0.0:8080/`


Once you finish installing the entire container and can access the url, open another terminal and generate the migrations for the project

A detailed step-by-step description is:
```
make migrateTables
```
This will generate migrations for your database.



And finally you can test the API by running the following command

open other terminal and run the CLI command

```
make runCLI
```

---


A guide on how to install docker for Linux, Mac and Windows is available here: https://docs.docker.com/get-docker/

Disclaimer: These instructions were tested using a Linux operating system, for Windows we suggest you install bash for Windows: https://itsfoss.com/install-bash-on-windows/

go1.22.2 linux/amd64 was used to develop and test this challenge.

---

### You can run other commands to validate the state of the api:

* To run the tests:
```
make test
```
