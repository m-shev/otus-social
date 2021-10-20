import {UserGender} from '../../../types';
import {FieldItem, FieldType} from '../../../components/Field/types';
import {cityOptions, interestOptions} from './options';

export type CreateFormValues = {
    firstName: string;
    lastName: string;
    age: number;
    gender: UserGender | null;
    interests: string[];
    city: string;
    email: string;
    password: string;
    password2: string;
};

export const MIN_USER_AGE = 10;
export const MAX_USER_AGE = 200;

export const initialValues: CreateFormValues = {
    firstName: '',
    lastName: '',
    age: 0,
    city: 'Москва',
    email: '',
    gender: UserGender.Male,
    interests: [],
    password: '',
    password2: '',
};
export const fieldList: FieldItem<CreateFormValues>[] = [
    {
        id: 'firstName',
        title: 'Имя',
        required: true,
        type: 'input',
    },
    {
        id: 'lastName',
        title: 'Фамилия',
        required: true,
        type: 'input',
    },
    {
        id: 'gender',
        title: 'Пол',
        required: true,
        type: 'select',
        options: [
            {
                value: UserGender.Male,
                title: 'Мужской',
            },
            {
                value: UserGender.Female,
                title: 'Женский',
            },
        ],
    },
    {
        id: 'age',
        title: 'Возраст',
        required: true,
        type: 'number',
        min: MIN_USER_AGE,
        max: MAX_USER_AGE,
    },
    {
        id: 'interests',
        title: 'Интересы',
        type: FieldType.Checkbox,
        options: interestOptions,
        required: false,
    },
    {
        id: 'city',
        title: 'Город',
        type: FieldType.Select,
        options: cityOptions,
        required: true,
    },
    {
        id: 'email',
        title: 'Электронная почта',
        type: 'email',
        required: true,
    },
    {
        id: 'password',
        title: 'Пароль',
        type: 'password',
        required: true,
    },
    {
        id: 'password2',
        title: 'Повторите пароль',
        type: 'password',
        required: true,
    },
];
