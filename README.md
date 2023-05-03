# go-lambda-s3-copy-automation
This severless function was created to be able to copy files from one S3 bucket into a subfolder in another S3 bucket cross account. After a user placed a file into this Bucket, a S3 Event Notification for the respected file triggered a PUT operation for the new file that triggered this lambda. The Event Notification has a prefix that it is looking for, so that it will only be triggered if a file with the specified prefix is stored to this bucket. Since S3 native Replication Rules do not support copying files from S3 to a different S3 into its SUBFOLDER, I created this lambda.

Two policies are added, both handle lambda access rights. "policy-s3-destination-account.json" will grant the desired permission so that the lambda can run copy the object to destionation bucket, whereas "policy-s3-source-bucket.json" allows the lambda to run in the source account and get the desired file.


#how to run
After you cloned this repo, run <b>go mod init main.go</b>, then run <b>go mod tidy</b>.
Now you can run the script to build the binary with <b>sh build.sh</b>
