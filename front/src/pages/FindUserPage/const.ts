import {FindUserForm} from '../../types';
import {FieldItem} from '../../components/Field/types';

export type FindForm = Pick<FindUserForm, 'name' | 'surname'>;

export const initialValues: FindForm = {
    surname: '',
    name: '',
};

export const fieldList: FieldItem<FindForm>[] = [
    {
        id: 'name',
        title: 'Имя',
        type: 'input',
        required: false,
    },
    {
        id: 'surname',
        title: 'Фамилия',
        type: 'input',
        required: false,
    },
];
