import * as React from 'react';
import styles from './CreatePostPage.module.scss';
import {Header} from '../../components/Header';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useFormik} from 'formik';
import {LoadableContent, SmallDots} from '../../components/LoadableContent';

export type CreatePostPageProps = {};

export const CreatePostPage: React.FC<CreatePostPageProps> = () => {
    const formik = useFormik({
        initialValues,
        onSubmit: () => {},
    });

    return (
        <div className={styles.root}>
            <Header />

            <h2>Создание поста</h2>

            <form>
                {fieldList.map((field) => {
                    return (
                        <Field
                            key={field.id}
                            error={formik.errors[field.id]}
                            setFieldValue={formik.setFieldValue}
                            {...formik.getFieldProps(field.id)}
                            {...field}
                        />
                    );
                })}

                <LoadableContent isLoading={false} Loader={<SmallDots />}>
                    <button type="submit" className={styles.btn}>
                        Создать
                    </button>
                </LoadableContent>
            </form>
        </div>
    );
};
