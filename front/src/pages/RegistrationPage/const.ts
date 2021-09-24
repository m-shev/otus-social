import {UserGender} from '../../types';
import {FieldProps} from './Field';
import {NumberField, SelectField} from './Field/types';

type DefaultFiledItem = Pick<FieldProps, 'id' | 'title' | 'required' | 'type'>;

type SelectFieldItem = DefaultFiledItem &
    Pick<SelectField, 'options' | 'type'> & {
        type: 'select';
    };

type NumberFieldItem = DefaultFiledItem & Pick<NumberField, 'min' | 'max' | 'type'>;

type FieldItem = DefaultFiledItem | SelectFieldItem | NumberFieldItem;

export const MIN_USER_AGE = 10;
export const MAX_USER_AGE = 200;

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
        type: 'input',
        required: false,
    },
    {
        id: 'city',
        title: 'Город',
        type: 'input',
        required: false,
    },
    {
        id: 'email',
        title: 'Электронная почта',
        type: 'email',
        required: true,
    },
];
