# Sprint 4

## Project Board
<https://github.com/shenoy-anurag/simple-reddit/projects/7>

## Backend
User Stories for backend

[Userstories and Issues Closed](https://github.com/shenoy-anurag/simple-reddit/issues?q=is%3Aissue+is%3Aclosed+label%3Asprint4+label%3A%22User+Stories+-+BE%22)

For Sprint 4, backend developers focused on Subscription and Save-To-Profile APIs and Testing and Deployment. Go tests cases have been written for Saved Posts, CRUD for Profiles and Comments APIs.

### Accomplishments:

#### APIs
- Get Saved Posts
- Get Saved Comments
- Update Saved Comments
- Updated Saved Posts
- Get Community Subscibers
- Update Subscribers
- Get Subscriptions
- Update Subscriptions

#### API Demo Video
https://user-images.githubusercontent.com/19914954/164345268-cbf9c181-1c5f-47f8-8940-ffdb18aefb9d.mp4

#### Go Tests:
- Saved Posts
- Saved Comments
- Community Subscribers
- Profile Subscriptions
- Signup For Profile
- Get Profile
- Edit Profile
- Delete Profile

#### Testing Video:
https://user-images.githubusercontent.com/19914954/164343869-386aba29-215e-423f-83cf-304df1b9028b.mp4

#### Bugs:
- Request validation error for GetAllCommunities API
- Update Subscription API isn't affecting the subscribers count of the community #288
- Security risk, the saved Posts APIs are returning DBModels #340
- Saved Comments API are returning DB Models, Security risk #341
- No limit on number of votes on a Post by a single User #346
- Only user who created post could vote on it #352

## Frontend

### Tasks
User Stories for Frontend
- View Posts
- Edit Posts
- Rate Posts
- Save Posts
- Add Comments
- Delete Comments
- Users can subscribe to a community
- Users can now only create a post or community only while logged in
- Dynamic pages for each post and each community
- Support
- Terms and Conditions
- Privacy Policy
- Content Policy
- Mod Policy

### Testing
- Updated Cypress testing for new system changes

[Userstories and Issues Closed](https://github.com/shenoy-anurag/simple-reddit/issues?q=is%3Aissue+is%3Aclosed+label%3Asprint4+label%3A%22User+Stories+-+FE%22)

### Accomplishments:

For Sprint 4, frontend devs focused on:
- Posts
- Posts,Communities,Comments,Raitings
- Edit existing Posts
- Terms,Privacy,Content,Mod Policies
- Support Info

### Application in Action
https://user-images.githubusercontent.com/34638485/164382796-d8b04846-25c4-40eb-9aaf-e7e070dba6ae.mp4

#### Angular Unit Testing and Testing using Cypress:

Frontend Demo video and wiki: <https://github.com/shenoy-anurag/simple-reddit/wiki/Demo#sprint-4-frontend-demo>
