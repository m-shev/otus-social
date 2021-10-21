/* eslint-disable @typescript-eslint/no-var-requires */
import faker from 'faker/locale/ru.js';
import fetch from 'node-fetch';


const LOOKUP = 100000;

const options = ['музыка', 'фильмы', 'спорт', 'компьютерные игры', 'путешествия'];

const generate = async () => {
    for (let i = 0; i < LOOKUP; i++) {
        const gender = faker.datatype.number({min: 0, max: 1});
        const user = {
            name: faker.name.firstName(gender),
            surname: faker.name.lastName(gender),
            age: faker.datatype.number({
                min: 12,
                max: 70,
            }),
            city: faker.address.city("ru"),
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

generate();
