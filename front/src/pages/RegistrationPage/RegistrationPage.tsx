import { Form, Formik } from 'formik';
import React from 'react';
import styles from './RegistrationPage.module.scss';
import { UserGender } from '../../type';
import { Field } from './Field';

export type RegistrationPageProps = {};
export type InitialValues = {
    firstName: string;
    lastName: string;
    age: number | null;
    gender: UserGender | null;
    interests: string;
    city: string;
    email: string;
    password: string;
}

const initialValues: InitialValues = {
  firstName: '',
  lastName: '',
  age: null,
  city: '',
  email: '',
  gender: null,
  interests: '',
  password: ''
};

export const RegistrationPage: React.FC<RegistrationPageProps> = () => {
  return <div className={styles.root}>
        <Formik
            initialValues={initialValues}
            onSubmit={() => {}}
        >
            {() => (
                <Form>
                    <Field type="input" name="name" title="Имя" required/>

                    <button type="submit">Зарегистрироваться</button>
                </Form>
            )}
        </Formik>
    </div>;
};
