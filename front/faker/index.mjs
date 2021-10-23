/* eslint-disable @typescript-eslint/no-var-requires */
import fs from 'fs';
import faker from 'faker/locale/ru.js';
import fetch from 'node-fetch';
import avatars from './avatars.js'

const LOOKUP = 200;

const options = ['музыка', 'фильмы', 'спорт', 'компьютерные игры', 'путешествия'];

const readCity = () => {
    const cityJson = JSON.parse(fs.readFileSync('russian-cities.json', 'utf8'))
    return cityJson.map(city => city.name);
}

const collectAvatars = async () => {
    const collectionMen = [];
    const collectionWomen = [];

    for (let i = 0; i < LOOKUP; i++) {
        const resp = await fetch('https://randomuser.me/api/?results=50');
        const json = await resp.json();
        console.log(json.results.length);
        json.results.forEach(item => {
            item.gender === 'female' ? collectionWomen.push(item.picture.large) : collectionMen.push(item.picture.large);
        })
    }

    const uniqWomen = Array.from(new Set(collectionWomen));
    const uniqMen = Array.from(new Set(collectionMen));

    console.log(uniqWomen);
    console.log(uniqMen);
    const dataSet = {
        women: uniqWomen,
        men: uniqMen
    }

    const str = JSON.stringify(dataSet);
    fs.writeFileSync('avatars.json', str, {encoding: 'utf8'});
}

const generate = async () => {
    const cities = readCity();

    for (let i = 0; i < LOOKUP; i++) {
        const gender = faker.datatype.number({min: 0, max: 1});
        const user = {
            avatar: faker.random.arrayElement(gender === 0 ? avatars.men : avatars.women),
            name: faker.name.firstName(gender),
            surname: faker.name.lastName(gender),
            gender: gender === 0 ? 'male' : 'female',
            age: faker.datatype.number({
                min: 12,
                max: 70,
            }),
            city: faker.random.arrayElement(cities),
            interests: faker.random.arrayElements(options),
            email: faker.internet.email(),
            password: 'password',
        };

        // console.log(user);
        const resp = await fetch('http://localhost:3005/user/registration', {
            method: 'post',
            body: JSON.stringify(user),
            headers: {'Content-Type': 'application/json'},
        });

        if (resp.status !== 200) {
            console.log('error', await resp.text());
        } else {
            // console.log('user created successfully', user.name, user.id);
        }
    }
};

// await collectAvatars();

await generate();
// readCity()