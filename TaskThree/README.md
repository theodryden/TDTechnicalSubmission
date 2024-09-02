 # Project overview
This project enjoyed usage of multiple AWS services: RDS (Postgresql), EC2 Instance (Linux) , Lambda, API Gateway. I've simplified my explanation into sections for the core learningings from  EC2, Database, and Lambda.


# Setting up EC2 Instance
AWS -> Launch EC2 Instance -> Download Keypair -> SSH into the EC2 -> Download files onto the server (Python, Postgresql) -> Uploaded the startup_script.py
![Screenshot 2024-09-02 at 22 24 45](https://github.com/user-attachments/assets/26dd4f22-74b3-4f47-a417-ad5a09c37ef2)


# Setting up RDS Database (Postgresql)
Created a new database -> Connected the database Lambda -> tested the Database worked via SSH and created a new database and table
![Screenshot 2024-09-02 at 22 46 40](https://github.com/user-attachments/assets/8376796a-f094-4a53-b515-b6e21a57775b)


Setup a budget report to ensure I didn't exceed payments. Luckily because AWS Aurora(postgresql) is different to the regular Postgresql. Leading me to promptly close AWS Aurora (Postgresql) DB and use normal Postgresql. 
![Screenshot 2024-09-02 at 22 26 43](https://github.com/user-attachments/assets/f807f62d-5606-4f08-96f9-0715ca85814b)


# Setting up AWS Lambda 
Created Lambda function for registration -> connected with other project components -> Setup Environment Variables -> tried coding logic in Python and JavaScript -> re-attemped using Layers .zip instead of code editor for importing modules -> found success in sending requests but received timedOut error with the database.

Attempted to resolve timeout requests by extending timeout to 2 minutes, checked RDS and EC2 were under the same VPC, tried different modules with Python, tried JavaScript.
![Screenshot 2024-09-02 at 22 40 42](https://github.com/user-attachments/assets/8ba3a897-9b30-4e8c-942b-b3ad96265910)
![Screenshot 2024-09-02 at 22 43 21](https://github.com/user-attachments/assets/bb33730c-1dfb-4f6f-b2c2-a84dadafe2b8)


# Future Improvement
* Resolve the network configuration issues causing a connection timeout with the AWS RDS Database and Lambda
* Review and test the API routes created on Gateway with Postman

