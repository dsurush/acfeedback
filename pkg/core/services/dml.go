package services

const userSaveDml = `Insert into 
"feedbacks" (feedBack, userId_Who, userId_Whom) 
values($1, $2, $3);`

const getFeedbackByIdDml  = `Select id, feedback, userid_who, userid_whom, remove from feedbacks where id = ($1)`

const removeFeedbackByIdDml = `update feedbacks set remove = false where id = ($1);`

const getFeedBackListDml = `Select id, feedback, userid_who, userid_whom, remove from feedbacks where remove = false;`