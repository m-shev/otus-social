import React from 'react';
import styles from './Field.module.scss';
import {FieldProps, FieldType} from './types';
import {SpecificField} from './SpecificField';

const renderCommonLabel = (props: Omit<FieldProps, 'setFieldValue'>) => {
    return props.type !== FieldType.Checkbox;
};

export const Field: React.FC<FieldProps> = (props) => {
    const {error, setFieldValue, ...restProps} = props;

    return (
        <div className={styles.root}>
            {renderCommonLabel(restProps) && (
                <label htmlFor={restProps.id} className={styles.label}>
                    {props.title}
                </label>
            )}

            <SpecificField {...restProps} setFieldValue={setFieldValue} />
            {error && <span className={styles.error}>{error}</span>}
        </div>
    );
};
