package sql_queries

const (
	InsertUser = `INSERT INTO users(about, email, fullname, nickname)
				    VALUES($1,$2,$3,$4);`

	UpdateUser = `UPDATE users SET about = $1, email = $2, fullname = $3
					WHERE LOWER(nickname) = LOWER($4);`

	SelectUserByNickname = `SELECT u.about, u.email, u.fullname, u.nickname 
							  FROM users as u
							  WHERE LOWER(u.nickname) = LOWER($1);`

	SelectUserByEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
						   FROM users as u
						   WHERE LOWER(u.email) = LOWER($1);`

	SelectUsersByNicknameAndEmail = `SELECT u.about, u.email, u.fullname, u.nickname 
									   FROM users as u
									   WHERE LOWER(u.nickname) = LOWER($1)
									   OR LOWER(u.email) = LOWER($2);`

	Truncate = `TRUNCATE ForumPosts;
				TRUNCATE UsersInForum;
				TRUNCATE TABLE votes CASCADE;
				TRUNCATE TABLE posts CASCADE;
				TRUNCATE TABLE threads CASCADE;
				TRUNCATE TABLE forums CASCADE;
				TRUNCATE TABLE users CASCADE;`

	SelectAll = `SELECT * FROM (SELECT COUNT(Posts.id) AS post FROM Posts) AS Post,
				(SELECT COUNT(Threads.id) AS thread FROM Threads) AS Thread,
				(SELECT COUNT(Forums.slug) AS forum FROM Forums) AS Forum,
				(SELECT COUNT(Users.nickname) AS user FROM Users) AS Users;`

	InsertForum = `INSERT INTO forums(slug, title, user_nick)
					 VALUES($1,$2,$3);`

	SelectForum = `SELECT posts, slug, threads, title, user_nick
					 FROM forums
					 WHERE LOWER(slug) = LOWER($1);`

	InsertThread = `INSERT INTO threads (author, forum, message, created, title, slug)
						VALUES ($1, (SELECT slug FROM forums WHERE LOWER(slug) = LOWER($2)), $3,
						COALESCE($4, NOW())::timestamptz, $5, NULLIF($6, ''))
						RETURNING id, slug, 0, forum`

	SelectThreadBySlug = `SELECT slug, title, message, forum, author, created, votes, id
							FROM threads 
							WHERE LOWER(slug) = LOWER($1)`

	SelectThreadById = `SELECT slug, title, message, forum, author, created, votes, id
							FROM threads 
							WHERE id=$1`

	QueryTemplateGetForumUsers = `SELECT about, email, fullname, nickname FROM users
									JOIN (SELECT nickname FROM UsersInForum WHERE forum=$1
										{{.Since}} ORDER BY nickname {{.Desc}} {{.Limit}}) as l
										USING (nickname) ORDER BY nickname {{.Desc}}`

	SelectThreadsByForum = `SELECT author, forum, created, id, message, slug, title, votes
								FROM threads 
								WHERE LOWER(forum) = LOWER($1) %s ORDER BY created %s %s`

	SelectThreadIdBySlug = `SELECT id FROM threads WHERE lower(slug)=lower($1)`

	SelectThreadIdById = `SELECT id FROM threads WHERE id=$1`

	InsertVote = `INSERT INTO votes(thread, author, vote) VALUES ($1, $2, $3)
					ON CONFLICT ON CONSTRAINT votes_thread_author_key DO
					UPDATE SET vote=$3 WHERE votes.thread=$1 AND lower(votes.author)=lower($2)`

	SelectPostById = `SELECT author, created, forum, id, message, thread, isedited, parent 
						FROM posts 
						WHERE id = $1`
)
