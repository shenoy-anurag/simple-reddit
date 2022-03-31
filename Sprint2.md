# Sprint 2
## Backend
User Stories for backend

[Userstories and Issues Closed](https://github.com/shenoy-anurag/simple-reddit/issues?q=is%3Aissue+is%3Aclosed+label%3Asprint2+label%3A%22User+Stories+-+BE%22)

The backend has taken up user stories pertaining to communities, profile and posts
So now the API to create, edit, get, delete posts, communities and profile has been created. Go tests cases are created for the posts, signup and login for the sprint.
CRUD tasks for posts, communities and profile has been created with exception of DELETE Profile.

### Accomplishments:

#### Communities
- Create Community
- Get Community Details
- Edit Community Details
- Delete Community Details

#### Posts
- Create Post
- Get Post
- Get All Posts
- Edit Post
- Delete Post

#### Profile
- Create Profile
- Get Profile Details
##### Note: DeleteProfile is moved to Sprint 3 as it depends on token authorization apart from username match.

#### Users
- Signup
- Login
- Check-Username

#### Go Tests:
- Added testing and unit-testing (using testing, testify, and net/http packages)
- Testing: User module (Signup, Login), Posts (Create, Delete, Edit, Get)

Checkout the test results [here](https://github.com/shenoy-anurag/simple-reddit/wiki/Demo#testing).

#### Bugs:
- Signup allows duplicate users
- Community creation allows duplicates


#### Bugs
- Signup allows duplicate users
- Community creation allows duplicates

Demo video and wiki: <https://github.com/shenoy-anurag/simple-reddit/wiki/Demo#sprint-2-back-end-demo>

## Frontend
User Stories for Frontend
- Added Unit testing and Cypress frontend testing
- Added Delete User Frontend
- Added Create Post Functionality
- Added Create Community Functionality
- Added Delete Subreddit Functionality
- Fixed Sign Up forms to dynamically change with user input
- SignUp now verifies username with database
- Connected Frontend and Backend processes with APIs
[Userstories and Issues Closed](https://github.com/shenoy-anurag/simple-reddit/issues?q=is%3Aissue+is%3Aclosed+label%3Asprint2+label%3A%22User+Stories+-+FE%22)

### Accomplishments:


#### Testing using Cypress:
- Navigation Menu Routing
- Proper New User SignUp
- Proper User Login
- Proper User Logout

Demo video and wiki: <https://github.com/shenoy-anurag/simple-reddit/wiki/Demo#sprint-2-frontend-testing-using-cypress>
