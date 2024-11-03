# aws-code-build-local-setup
This repository includes an example to setup AWS codebuild locally

Use this file to check how we can utilize it setup codebuild to build binaries locally or to create a docker image and push that to ECR 

Command to build local binaries: 


```
./codebuild_build.sh -i <build_image>:<build_image_tag> -a <output_directory> -s <source_directory>  -l <local_code_build_agent_image>:<aarch64/latest> 
```

Command to build and Push docker image : use `buildspec-docker.yml` from the repository:

```
./codebuild_build.sh -i <build_image>:<build_image_tag> -a <output_directory> -s <source_directory>  -l <local_code_build_agent_image>:<aarch64/latest> -b buildspec-docker.yml -d -e envfile -c ~/.aws/credentials  -p default
```


Usage : (Copied from codebuild_build.sh ) (All Rights reserved with the original author of the script.)

codebuild_build.sh [-i image_name] [-a artifact_output_directory] [options]
Required:
  -i        Used to specify the customer build container image.
  -a        Used to specify an artifact output directory.
Options:
  -l IMAGE  Used to override the default local agent image.
  -r        Used to specify a report output directory.
  -s        Used to specify source information. Defaults to the current working directory for primary source.
               * First (-s) is for primary source
               * Use additional (-s) in <sourceIdentifier>:<sourceLocation> format for secondary source
               * For sourceIdentifier, use a value that is fewer than 128 characters and contains only alphanumeric characters and underscores
  -c        Use the AWS configuration and credentials from your local host. This includes ~/.aws and any AWS_* environment variables.
  -p        Used to specify the AWS CLI Profile.
  -b FILE   Used to specify a buildspec override file. Defaults to buildspec.yml in the source directory.
  -m        Used to mount the source directory to the customer build container directly.
  -d        Used to run the build container in docker privileged mode.
  -e FILE   Used to specify a file containing environment variables.
            (-e) File format expectations:
               * Each line is in VAR=VAL format
               * Lines beginning with # are processed as comments and ignored
               * Blank lines are ignored
               * File can be of type .env or .txt
               * There is no special handling of quotation marks, meaning they will be part of the VAL
               