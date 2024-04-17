# The Command Line Interface

The best code is executing code and code is not executed in a vacuum. It usually has entrances and exits.

We have created a CLI in the console so that you can test the API, and validate that it saves the information in the db, and sends the email with the summary of the information

Your code will be called invoking the file.

* transaction_info.csv

Here are some examples:

```bash
./cli transaction_info.csv
```
This will send the information from the csv file to the api, where it will process the information, store it in the db for data persistence, and send the email with the data that the exercise required