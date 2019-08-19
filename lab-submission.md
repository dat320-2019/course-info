# Autograder

This course uses Autograder, a tool developed at the University of Stavanger
for students and teaching staff to manage the submission and validation of
lab assignments. All lab submissions from students are handled using Git,
a source code management system, and GitHub, a web-based hosting service for
Git source repositories. Thus, basic knowledge of these tools are necessary.
The procedure used to submit your lab assignments is explained below.

Students push their updated lab submissions to GitHub. Every lab submission is
then processed by a custom continuous integration tool. This tool will run
several test cases on the submitted code. Autograder generates feedback that
let the students verify if their submission implements the required functionality.
This feedback is available through a web interface. The feedback from the
Autograder system can be used by students to improve their submissions.

## Git and GitHub

Git is a distributed revision control and source code management system. Basic
knowledge of Git is required for handing in the lab assignments. There are many
resources available online for learning Git. A good book is Pro Git by Scott
Chacon. The book is available for free [here](https://git-scm.com/book).
Chapter 2.1 and 2.2 should contain the necessary information for delivering the
lab assignments.

GitHub is a web-based hosting service for software development projects that
use the Git revision control system. An introduction to Git and GitHub is
available in [this video](http://youtu.be/U8GBXvdmHT4). Students need to sign
up for a GitHub account to get access to the needed course material.

## Autograder Registration

Follow the steps below to register and sign up for the course on Autograder.

1. Go to [GitHub](http://github.com) and register. A GitHub account is required
   to sign in to Autograder. You can skip this step if you already have an
   account.

2. Click the "Sign in with GitHub" button in
   [Autograder](http://ag.itest.run) to register. You will then be
   taken to GitHub's website.

3. Approve that our Autograder application may have permission to access to the
   requested parts of your account. It is possible to make a separate GitHub
   account for only this (and other) courses if you do not want Autograder to
   access your personal one with the requested permissions.

## Signing up for the course

1. Click the Courses menu item.

2. In the Courses menu select “Join a course”. Available courses will be listed.

3. Find the course “dat320-YEAR” and click Sign up.

4. Read through and accept the terms. You will then be invited to the
   organization "dat320-YEAR" on GitHub.

5. An invitation will be sent to the email address registered with your GitHub
   account. Accept the invitation using the received email.

6. Wait for the teaching staff to verify your Autograder-registration.

7. You will get your own repository in the "dat320-YEAR" organization on GitHub
   after your registration is verified.

## Instructions for submitting a lab assignment

### TODO(Meling) Fix instructions

This section offers step-by-step instructions on how to hand in
Lab 1. Please refer to the workflow described below also for future labs unless
otherwise noted.

1. You will have access to two repositories when you have signed up on
   Autograder. The first is the `labs` repository, which is where we will
   publish all lab assignments, skeleton code and additional information
   needed. You only have read access to this repository. The second repository
   is your own repository named `username-labs`, where `username` is your
   GitHub username. You have write access to this repository. Your answers to
   the assignments should be pushed here.

2. There are several ways to get the `labs` repo. Below we describe two options.

3. (Option 1) Using just the `git` command on the command line.
    1. `ls -la $GOPATH` (make sure that this is a real directory)
    2. `mkdir -p $GOPATH/src/github.com/dat320-2019`
    3. `cd $GOPATH/src/github.com/dat320-2019`
    4. `git clone https://github.com/dat320-2019/assignments.git`
    5. You will be asked for your GitHub user name and password
    6. Continue in Step 6 below

4. (Option 2) Using the `go get` command on the command line.
    1. `ls -la $GOPATH` (make sure that this is a real directory)
    2. `go get github.com/dat320-2019/assignments`
    3. You will be asked for your GitHub user name and password
    4. (ignore the warning message about no buildable Go files)
    5. Continue in Step 6 below

5. (Advanced Option) Advanced users: Follow these [steps](https://github.com/dat320-2019/course-info/blob/master/github-ssh.md) if you want to use SSH for GitHub authentication. This will help avoid having to type your password every time.

6. (Information) Both Option 1 or 2 above will clone the original `labs` git
   repo (not your copy of it.) This means that you don't need to
   change the import path in the source files to use your own repository's
   path. That is, when you make a commit and push to submit your handin, you
   don't have to change this back to the original import path.

7. Now we need to set up your own remote so that you can make changes and push
   those changes to your own copy of the repo. Follow these instructions:
    1. `cd $GOPATH/src/github.com/dat320-2019/assignments`.
    2. `git remote add labs https://github.com/dat320-2019/username-labs`
     where `username` should be replaced with your own GitHub username.
    3. The above command adds your own `username-labs` repository as a remote
   repository on your local machine. This means that once you've modified some
   files and committed the changes locally, you can run:
    4. `git push labs` to have them pushed up to your own `username-labs` repository on GitHub.

8. If you make changes to your own `username-labs` repository using the GitHub
   web interface, and want to pull those changes down to your own computer, you
   can run the command:
    * `git pull labs master`
    * In later labs, you will work in groups. This approach is also the way that you can download (pull) your group's code changes from GitHub, assuming that another group member has previously pushed it out to GitHub.

9. As time goes by we (the teaching staff) may publish updates to the
   original `labs` repo, e.g. new or updated lab assignments. To see these
   updates, you will need to run the following command:
    * `git pull origin master`.

10. For the first lab, you will submit the source code for the hello world
   application and the answers to the three sets of questions. The skeleton code
   for the application, `hello.c`, has been provided along with forms for the
   questions. Autograder will run a set of test cases to verify your
   implementation. Not all tests must pass to get a passing grade.

11. In the following, we will use `hello.c` as an example. Change directory to:
   `cd $GOPATH/src/github.com/dat320-2019/assignments/lab1` and confirm that the files
   for lab1 resides in that folder. They should, assuming that you ran the `go
   get` command earlier.

12. Implement the main function in `hello.c`.

13. When your application is working, you may push your code to GitHub. This will
    trigger Autograder which will then run a test suite on your code.

14. Using `hello.c` as an example, use the following
    procedure to commit and push your changes to GitHub and Autograder:

    ```console
    $ cd $GOPATH/src/github.com/dat320-2019/assignments/lab1
    $ git add hello.c
    $ git commit
    // This will open an editor for you to write a commit message
    // Use for example "Implemented Assignment 1"
    $ git push labs
    ```

15. Running the last command above will, due to an error on our part, result in
    Git printing an error message about a conflict between the `README.md` file
    in the `labs` repository and the `README.md` file in your `username-labs`
    repository. Here is how to fix it:

    ```console
    $ git push labs
    ...
    ! [rejected]        master -> master (fetch first)
    error: failed to push some refs to 'git@github.com:dat320-2019/username-labs.git'
    ...
    $ git pull labs master
    ...
    Auto-merging README.md
    CONFLICT (add/add): Merge conflict in README.md
    Automatic merge failed; fix conflicts and then commit the result.
    ...
    $ cd $GOPATH/src/github.com/dat320-2019/assignments
    $ vi README.md
    // Remove everything in the file, then add for example "username-labs" to the file.
    // Save and exit.
    $ git add README.md
    $ git commit
    $ // Use the existing (merge) commit message. Save and exit.
    $ git push labs
    // Your push should now complete successfully.
    // You may check that your changes are reflected on GitHub through the GitHub web interface.
    //If you still get an error like "fatal: Exiting because of an unresolved conflict." use the following command
    $ git pull --allow-unrelated-histories labs master
    $ git add README.md
    $ git commit
    $ git push labs
    ```

16. Autograder will now build and run a test suite on the code you submitted.
    You can check the output by going the [Autograder web
    interface](http://ag.itest.run/). The results (build log) should be
    available under "Individual - lab1". Note that the results shows output
    for all the tests in current lab assignment. You will want to focus on the
    output for the specific test results related to the task you're working on.

17. Repeat step 15 for the three sets of answers.

18. If some of the autograder tests fail, you may make changes to your code/answers.

19. Push your changes using `git push labs`. You should be able to view your
    results in the Autograder web interface as described earlier.

## Lab Approval

To have your lab assignment approved, you must come to the lab during lab hours
and present your solution. This lets you present the thought process behind your
solution, and gives us a more information for grading purposes. When you are
ready to show your solution, reach out to a member of the teaching staff.
It is expected that you can explain your code and show how it works.
You may show your solution on a lab workstation or your own
computer. The results from Autograder will also be taken into consideration
when approving a lab. At least 60% of the Autograder tests should pass for the
lab to be approved. A lab needs to be approved before Autograder will provide
feedback on the next lab assignment.

Also see the [Grading and Collaboration
Policy](https://github.com/dat320-2019/course-info/blob/master/policy.md)
document for additional information.
