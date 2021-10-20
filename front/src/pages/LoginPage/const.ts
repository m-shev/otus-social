import {LoginForm} from '../../types';
import {FieldItem} from '../../components/Field/types';

export const initialValues: LoginForm = {
    login: '',
    password: '',
};

export const fieldList: FieldItem<LoginForm>[] = [
    {
        id: 'login',
        title: 'Логин',
        required: true,
        type: 'input',
    },
    {
        id: 'password',
        title: 'Пароль',
        required: true,
        type: 'password',
    },
];
