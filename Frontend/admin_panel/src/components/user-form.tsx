'use client';
import { Formik, Form, Field, ErrorMessage } from "formik";
import { userSchema, UserFormValues } from "@/lib/validation";
import { Input } from "@/components/ui/input";
import { Button } from "./ui/button";
import { Label } from "@/components/ui/label";

export const UserForm = ({
    defaultValues,
    onSubmit,
    submitLabel = 'Save',
}: {
    defaultValues?: Partial<UserFormValues>;
    onSubmit: (data: UserFormValues) => void;
    submitLabel?: string;
}) => {
    const initialValues: UserFormValues = {
        firstName: '',
        lastName: '',
        email: '',
        role: 'Viewer',
        isActive: true,
        ...defaultValues
    };

    return (
        <Formik
            initialValues={initialValues}
            validationSchema={userSchema}
            onSubmit={async (values, { setSubmitting }) => {
                await onSubmit(values);
                setSubmitting(false);
            }}
            enableReinitialize
        >
            {({ isSubmitting }) => (
                <Form className="space-y-4">
                    <div>
                        <Label htmlFor="firstName">First Name</Label>
                        <Field as={Input} id="firstName" name="firstName" />
                        <ErrorMessage name="firstName" component="p" className="text-sm text-red-600 mt-1" />
                    </div>
                    <div>
                        <Label htmlFor="lastName">Last Name</Label>
                        <Field as={Input} id="lastName" name="lastName" />
                        <ErrorMessage name="lastName" component="p" className="text-sm text-red-600 mt-1" />
                    </div>
                    <div>
                        <Label htmlFor="email">Email</Label>
                        <Field as={Input} id="email" type="email" name="email" />
                        <ErrorMessage name="email" component="p" className="text-sm text-red-600 mt-1" />
                    </div>
                    <div>
                        <Label htmlFor="role">Role</Label>
                        <Field as="select" id="role" name="role" className="mt-1 block w-full rounded-md border border-gray-300 bg-white px-3 py-2 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
                            <option value="Admin">Admin</option>
                            <option value="Editor">Editor</option>
                            <option value="Viewer">Viewer</option>
                        </Field>
                        <ErrorMessage name="role" component="p" className="text-sm text-red-600 mt-1" />
                    </div>
                    <div className="flex items-center space-x-2">
                        <Field type="checkbox" id="isActive" name="isActive" className="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded" />
                        <Label htmlFor="isActive" className="mb-0">Active</Label>
                    </div>
                    <ErrorMessage name="isActive" component="p" className="text-sm text-red-600 mt-1" />
                    <Button type="submit" disabled={isSubmitting}>{submitLabel}</Button>
                </Form>
            )}
        </Formik>
    );
}


// ---------------------React Hook Form---------------------------------

// 'use client';
// import { useEffect } from "react";
// import { useForm } from "react-hook-form";
// import { yupResolver } from "@hookform/resolvers/yup";
// import { userSchema, UserFormValues } from "@/lib/validation";
// import { Input } from "@/components/ui/input";
// import { Button } from "./ui/button";
// import { Label } from "@/components/ui/label";

// export const UserForm = ({
//     defaultValues,
//     onSubmit,
//     submitLabel = 'Save',
// }: {
//     defaultValues?: Partial<UserFormValues>;
//     onSubmit: (data: UserFormValues) => void;
//     submitLabel?: string;
// }) => {
//     const { register, handleSubmit, formState: { errors, isSubmitting }, reset } = useForm<UserFormValues>({
//         resolver: yupResolver(userSchema),
//         defaultValues: {
//             firstName: '',
//             lastName: '',
//             email: '',
//             role: 'Viewer',
//             isActive: true,
//             ...defaultValues
//         }
//     });

//     useEffect(() => {
//         reset({
//             firstName: '',
//             lastName: '',
//             email: '',
//             role: 'Viewer',
//             isActive: true,
//             ...defaultValues
//         });
//     }, [defaultValues, reset]);

//     return (
//         <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
//             <div>
//                 <Label htmlFor="firstName">First Name</Label>
//                 <Input id="firstName" {...register('firstName')} />
//                 {errors.firstName && <p className="text-sm text-red-600 mt-1">{errors.firstName.message}</p>}
//             </div>
//             <div>
//                 <Label htmlFor="lastName">Last Name</Label>
//                 <Input id="lastName" {...register('lastName')} />
//                 {errors.lastName && <p className="text-sm text-red-600 mt-1">{errors.lastName.message}</p>}
//             </div>
//             <div>
//                 <Label htmlFor="email">Email</Label>
//                 <Input id="email" type="email" {...register('email')} />
//                 {errors.email && <p className="text-sm text-red-600 mt-1">{errors.email.message}</p>}
//             </div>
//             <div>
//                 <Label htmlFor="role">Role</Label>
//                 <select id="role" {...register('role')} className="mt-1 block w-full rounded-md border border-gray-300 bg-white px-3 py-2 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm">
//                     <option value="Admin">Admin</option>
//                     <option value="Editor">Editor</option>
//                     <option value="Viewer">Viewer</option>
//                 </select>
//                 {errors.role && <p className="text-sm text-red-600 mt-1">{errors.role.message}</p>}
//             </div>
//             <div className="flex items-center space-x-2">
//                 <input id="isActive" type="checkbox" {...register('isActive')} className="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded" />
//                 <Label htmlFor="isActive" className="mb-0">Active</Label>
//             </div>
//             {errors.isActive && <p className="text-sm text-red-600 mt-1">{errors.isActive.message}</p>}
//             <Button type="submit" disabled={isSubmitting}>{submitLabel}</Button>
//         </form>
//     );
// }
