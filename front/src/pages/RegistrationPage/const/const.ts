import {UserGender} from '../../../types';
import {
    CheckboxFieldProps,
    FieldProps,
    FieldType,
    NumberFieldProps,
    SelectFieldProps,
} from '../Field/types';
import {cityOptions, interestOptions} from './options';

type DefaultFiledItem = Pick<FieldProps, 'id' | 'title' | 'required' | 'type'>;

type SelectFieldItem = DefaultFiledItem &
    Pick<SelectFieldProps, 'options' | 'type'> & {
        type: 'select';
    };

type NumberFieldItem = DefaultFiledItem & Pick<NumberFieldProps, 'min' | 'max' | 'type'>;

type FieldItem = DefaultFiledItem | SelectFieldItem | NumberFieldItem | CheckboxFieldProps;

export type InitialValues = {
    firstName: string;
    lastName: string;
    age: number | undefined;
    gender: UserGender | null;
    interests: string;
    city: string;
    email: string;
    password: string;
};

export const MIN_USER_AGE = 10;
export const MAX_USER_AGE = 200;

export const initialValues: InitialValues = {
    firstName: '',
    lastName: '',
    age: undefined,
    city: '',
    email: '',
    gender: UserGender.Male,
    interests: '',
    password: '',
};
export const fieldList: FieldItem[] = [
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
];