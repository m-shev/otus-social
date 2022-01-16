import * as React from 'react';
import styles from './Field.module.scss';
import {TextAreaFieldProps} from './types';

export type TextAreaProps = TextAreaFieldProps;

export const TextArea: React.FC<TextAreaProps> = (props) => {
    return <textarea className={`${styles.textarea} ${styles.field}`} {...props} />;
};
