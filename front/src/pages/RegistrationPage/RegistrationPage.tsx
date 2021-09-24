import React from 'react';
import {useFormik} from 'formik';
import styles from './RegistrationPage.module.scss';
import {UserGender} from '../../types';
import {fieldList} from './const';
import {Field} from './Field';

export type RegistrationPageProps = Record<string, never>;
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

const initialValues: InitialValues = {
    firstName: '',
    lastName: '',
    age: undefined,
    city: '',
    email: '',
    gender: UserGender.Male,
    interests: '',
    password: '',
};

export const RegistrationPage: React.FC<RegistrationPageProps> = () => {
    const formik = useFormik({
        initialValues,
        onSubmit: () => {},
    });

    return (
        <div className={styles.root}>
            <form className={styles.form} onSubmit={formik.handleSubmit}>
                <div className={styles.fields}>
                    {fieldList.map((field) => {
                        return (
                            <Field key={field.id} {...field} {...formik.getFieldProps(field.id)} />
                        );
                    })}
                </div>

                <button type="submit" className={styles.submit}>
                    Зарегистрироваться
                </button>
            </form>
        </div>
    );
};
