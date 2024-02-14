-- Insertion de données fictives dans la table User
INSERT INTO User (username, email, `password`)
VALUES (
        'AliceSmith',
        'alice.smith@example.com',
        'MotDePasse123'
    ),
    (
        'BobJohnson',
        'bob.johnson@example.com',
        'Secret456'
    ),
    (
        'CharlieBrown',
        'charlie.brown@example.com',
        'P@ssw0rd'
    ),
    (
        'DavidMiller',
        'david.miller@example.com',
        'Secure789'
    ),
    (
        'EmmaWhite',
        'emma.white@example.com',
        'Confidential987'
    );
-- Insertion de données fictives dans la table Category
INSERT INTO Category (name)
VALUES ('Religion'),
    ('Technologie'),
    ('Voyages'),
    ('Cuisine'),
    ('Sport'),
    ('Mode');
-- Insertion de données fictives dans la table Post
INSERT INTO Post (content, creation_date, `user_id`)
VALUES (
        'Nouvelle découverte technologique !',
        '2022-02-10 13:02:25',
        1
    ),
    (
        'Mes vacances à la plage de la Guinee',
        '2022-02-11 01:56:00',
        2
    ),
    (
        'Recette secrète de grand-mère vraiment secrete hein',
        '2022-02-12 21:23:47',
        3
    ),
    (
        'Entraînement intensif maintenant',
        '2022-02-13 07:26:12',
        4
    ),
    (
        'Tendances de la mode automne 2022',
        '2022-02-14 20:01:30',
        5
    );
-- Insertion de données fictives dans la table Comment
INSERT INTO Comment (comment, date_creation, post_id, `user_id`)
VALUES ('Super article !', '2022-02-14 20:02:30', 1, 3),
    ("J'adore cet endroit ! ", '2022-03-14 10:02:30', 5, 3),
    ("J'me demande pourquoi autant de beauté ! ", '2022-02-15 06:02:30', 5, 3),
    (' Délicieux, merci pour la recette ! ', '2022-02-16 16:10:30', 3, 3),
    (' Continuez comme ça ! ', '2022-02-15 16:12:50', 4, 1),
    (' Haha mdr c est dur actuellement fori ! ', '2022-03-14 21:42:30', 4, 3),
    (' Vraiment je vous encourage de nouveau !! Continuez comme ça ! ', '2022-02-17 20:02:30', 4, 1),
    (' Les nouvelles tendances sont incroyables ', '2022-02-24 14:42:20', 5, 1);
-- Insertion de données fictives dans la table PostCategory
INSERT INTO PostCategory (post_id, category_id)
VALUES (1, 1),
    (2, 2),
    (2, 4),
    (3, 5),
    (3, 3),
    (4, 4),
    (4, 2),
    (4, 1),
    (5, 5);
-- Insertion de données fictives dans la table LikeDislike
INSERT INTO LikeDislike (user_id, post_id, liked, disliked)
VALUES (1, 1, true, false),
    (2, 2, false, true),
    (2, 3, true, false),
    (3, 1, true, false),
    (3, 4, false, true),
    (4, 4, true, false),
    (5, 4, true, false);