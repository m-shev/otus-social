import React from 'react';
import {useFormik} from 'formik';
import styles from './RegistrationPage.module.scss';
import {fieldList, initialValues} from './const';
import {Field} from './Field';

export type RegistrationPageProps = Record<string, never>;

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
