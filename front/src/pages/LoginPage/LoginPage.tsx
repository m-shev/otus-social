import * as React from 'react';
import {useFormik} from 'formik';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useOnLogin} from './hooks';
import styles from './LoginPage.module.scss';
import {useStore} from 'effector-react';
import {$userStore} from '../../store/user';

export type LoginPageProps = {};

export const LoginPage: React.FC<LoginPageProps> = () => {
    const {onSubmit, isFetch, requestState, error} = useOnLogin();
    const {user} = useStore($userStore);

    const formik = useFormik({
        initialValues,
        onSubmit: onSubmit,
    });

    return (
        <div>
            <h1>Авторизация</h1>

            {error && (
                <div className={styles.error}>{`Не удалось авторизоваться: ${error.message}`}</div>
            )}

            {requestState !== 'success' ? (
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
            ) : (
                <span>{`Добро пожаловать ${user?.name}`}</span>
            )}
        </div>
    );
};
