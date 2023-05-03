# go-lambda-s3-copy-automation
This severless function was created to be able to copy files from one S3 bucket into a subfolder in another S3 bucket cross account. After a user placed a file into this Bucket, a S3 Event Notification for the respected file triggered a PUT operation for the new file that triggered this lambda. Since S3 native Replication Rules do not support copying files from S3 to a different S3 into its SUBFOLDER, I created this lambda.

