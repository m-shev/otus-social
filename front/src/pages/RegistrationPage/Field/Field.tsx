import React from 'react';
import { Field as FormikField } from 'formik';
import styles from './Field.module.scss';
export type FieldProps = {
    name: string;
    title: string;
    type: string;
    required: boolean;
};

export const Field: React.FC<FieldProps> = ({ name, title, type, required }) => {
  return (
        <div className={styles.root}>
            <label htmlFor={name}>{title}</label>
            <FormikField id={name} name={name} type={type} required />
        </div>
  );
};
