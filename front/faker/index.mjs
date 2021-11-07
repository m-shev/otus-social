/* eslint-disable @typescript-eslint/no-var-requires */
import fs from 'fs';
import faker from 'faker/locale/ru.js';
import fetch from 'node-fetch';
import avatars from './avatars.js'

const LOOKUP = 10_000;

const options = ['музыка', 'фильмы', 'спорт', 'компьютерные игры', 'путешествия'];
const localUrl = 'http://localhost:3005/user/registrations';
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
    let count = 0;
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


        // const resp = fetch('https://mshev.ru/user/registration', {
        //     method: 'post',
        //     body: JSON.stringify(user),
        //     headers: {'Content-Type': 'application/json'},
        // });

        users.push(user);

        if (users.length === 500) {
            try {
                console.time("executed")
                const resp = await fetch(localUrl, {
                    method: 'post',
                    body: JSON.stringify(users),
                    headers: {'Content-Type': 'application/json'}
                })

                if (resp.status !== 200) {
                    console.log(await resp.text())
                }

                count += users.length;
                console.log("success added: ", count);
                console.timeEnd("executed");
            } catch (e) {
                console.log("error", e);
            }

            users = [];
        }
    }

    console.log("finished")
};

// await collectAvatars();

await generate();
// readCity()