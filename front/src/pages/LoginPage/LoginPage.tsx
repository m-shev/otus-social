import * as React from 'react';
import {useFormik} from 'formik';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useOnLogin} from './hooks';
import styles from './LoginPage.module.scss';
import {useRedirectToProfile} from '../../hooks';
import {Header} from '../../components/Header';
import {useHistory} from 'react-router';

export type LoginPageProps = {};

export const LoginPage: React.FC<LoginPageProps> = () => {
    const {onSubmit, isFetch, error} = useOnLogin();
    const history = useHistory();
    useRedirectToProfile();

    const formik = useFormik({
        initialValues,
        onSubmit: onSubmit,
    });

    return (
        <div className={styles.root}>
            <Header showLoginButton={false} />
            <h2>Авторизация</h2>

            <div className={styles.content}>
                {error && (
                    <div
                        className={styles.error}
                    >{`Не удалось авторизоваться: ${error.message}`}</div>
                )}

                <form onSubmit={formik.handleSubmit} className={styles.form}>
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

                    {!isFetch && (
                        <button
                            onClick={() => {
                                return history.push('/registration');
                            }}
                            className={styles.submit}
                        >
                            Зарегистрироваться
                        </button>
                    )}
                </form>
            </div>
        </div>
    );
};
