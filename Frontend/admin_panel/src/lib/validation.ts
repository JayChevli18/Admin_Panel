import * as yup from 'yup';

export const userSchema = yup.object().shape({
    firstName: yup.string().required('First name is required').min(2, 'First name must be at least 2 characters'),
    lastName: yup.string().required('Last name is required').min(2, 'Last name must be at least 2 characters'),
    email: yup.string().email('Invalid email format').required('Email is required'),
    role: yup.mixed<'Admin' | 'Editor' | 'Viewer'>().oneOf(['Admin', 'Editor', 'Viewer']).required('Role is required'),
    isActive: yup.boolean().required('Active status is required'),
});

export type UserFormValues = yup.InferType<typeof userSchema>;