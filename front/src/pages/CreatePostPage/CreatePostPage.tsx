import * as React from 'react';
import styles from './CreatePostPage.module.scss';
import {Header} from '../../components/Header';
import {fieldList, initialValues} from './const';
import {Field} from '../../components/Field';
import {useFormik} from 'formik';
import {LoadableContent, SmallDots} from '../../components/LoadableContent';
import {useCreatePost} from './hooks';
import {useRedirectToProfile} from '../../hooks';

export type CreatePostPageProps = {};

export const CreatePostPage: React.FC<CreatePostPageProps> = () => {
    const {onCreatePost, isFetch, error, requestState} = useCreatePost();

    const formik = useFormik({
        initialValues,
        onSubmit: onCreatePost,
    });

    useRedirectToProfile(() => {
        return requestState === 'success';
    });

    return (
        <div className={styles.root}>
            <Header />

            <h2>Создание поста</h2>

            <form onSubmit={formik.handleSubmit}>
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

                {error && <div className={styles.error}>{error.message}</div>}

                <LoadableContent isLoading={isFetch} Loader={<SmallDots />}>
                    <button type="submit" className={styles.btn}>
                        Создать
                    </button>
                </LoadableContent>
            </form>
        </div>
    );
};
