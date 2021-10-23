import React from 'react';
import {useFormik} from 'formik';
import styles from './RegistrationPage.module.scss';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useCreateUser} from './hooks';
import {useHistory} from 'react-router';
import {Header} from '../../components/Header';

const REDIRECT_DELAY = 2500;

export type RegistrationPageProps = Record<string, never>;

export const RegistrationPage: React.FC<RegistrationPageProps> = () => {
    const {onSubmit, isFetch, requestState, error} = useCreateUser();

    const formik = useFormik({
        initialValues,
        onSubmit: onSubmit,
    });

    const pushHistory = useHistory();

    if (requestState === 'success') {
        setTimeout(() => {
            pushHistory.push('/');
        }, REDIRECT_DELAY);
    }

    return (
        <div className={styles.root}>
            <Header />
            <h2>Регистрация</h2>
            {requestState !== 'success' ? (
                <form className={styles.form} onSubmit={formik.handleSubmit}>
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
                        {isFetch ? '...' : ' Зарегистрироваться'}
                    </button>
                </form>
            ) : (
                <>
                    <span>Поздравляем с успешной регистрацией</span>
                    <span>Сейчас вы будете направлены на главную страницу...</span>
                </>
            )}
        </div>
    );
};
