import * as React from 'react';
import {useFormik} from 'formik';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useOnLogin} from './useOnLogin';
import styles from './LoginPage.module.scss';

export type LoginPageProps = {};

export const LoginPage: React.FC<LoginPageProps> = () => {
    const {onSubmit, isFetch, requestState, error} = useOnLogin();

    const formik = useFormik({
        initialValues,
        onSubmit: onSubmit,
    });

    return (
        <div>
            <h1>Авторизация</h1>

            <form onSubmit={formik.handleSubmit}>
                <div className={styles.fields}>
                    {fieldList.map((field) => {
                        return (
                            <Field
                                error={formik.errors[field.id]}
                                key={field.id}
                                {...field}
                                {...formik.getFieldProps(field.id)}
                                setFieldValue={formik.setFieldValue}
                            />
                        );
                    })}
                </div>

                <button type="submit" className={styles.submit} disabled={isFetch}>
                    {isFetch ? '...' : 'Войти'}
                </button>
            </form>

            <button
                className={styles.submit}
                disabled={isFetch}
                onClick={() => {
                    fetch('http://localhost:3005/user/test', {credentials: 'include'});
                }}
            >
                {isFetch ? '...' : 'test'}
            </button>
        </div>
    );
};
