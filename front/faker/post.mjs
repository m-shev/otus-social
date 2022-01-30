import faker from 'faker/locale/ru.js';
import fetch from 'node-fetch';

const baseUrl = 'http://localhost:3005';

const userFriendsUrl = (userId) => `${baseUrl}/user/${userId}/friends?skip=0&take=1000`;
const createPostUrl = () => `${baseUrl}/post`;

const userIds = [11819]

const createPost = (number, authorId) => {
    const posts = [];
    for (let i = 0; i < number; i++) {
        posts.push({
            authorId,
            content: faker.lorem.text(),
            imageLink: faker.image.image(),
        })
    }
    return posts;
}

const generate = async () => {
    const resp = await fetch(userFriendsUrl(userIds[0]))

    if (resp.status !== 200) {
        throw new Error(await resp.text())
    }

    const friendList = (await resp.json()).friendList;

    friendList.forEach((friend) => {
        const promises = [];
        const posts = createPost(200, friend.id)
        posts.forEach(post => {
            promises.push(fetch(createPostUrl(), {
                method: 'post',
                body: JSON.stringify(post),
            }))
        })
        console.log(posts);
    })


}


await generate();