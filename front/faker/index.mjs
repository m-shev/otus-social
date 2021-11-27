/* eslint-disable @typescript-eslint/no-var-requires */
import fs from 'fs';
import faker from 'faker/locale/ru.js';
import fetch from 'node-fetch';
import avatars from './avatars.js'

const LOOKUP = 100_000;

const options = ['музыка', 'фильмы', 'спорт', 'компьютерные игры', 'путешествия'];
const localUrl = 'http://localhost:3005/user/registration';
const promUrl = 'https://mshev.ru/user/registrations';

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
    let users = [];

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
            email: i.toString() + faker.internet.email(),
            password: 'password',
        };


        users.push(user);

        if (users.length === 200) {
            console.time("executed");
            const promises = [];

            for (const user of users) {
                const resp = fetch(localUrl, {
                    method: 'post',
                    body: JSON.stringify(user),
                    headers: {'Content-Type': 'application/json'}
                })

                promises.push(resp);
            }

            const results = await Promise.all(promises).catch(e => {
                console.log(e);
            });

            for (const res of results) {
                res.status !== 200 ? console.log(await res.text()) : console.log(await res.json());
            }

            users = [];
            console.timeEnd("executed");
        }
    }

    console.log("finished")
};

console.time("total execution");
await generate();
console.timeEnd("total execution");
