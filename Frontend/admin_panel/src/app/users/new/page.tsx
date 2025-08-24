'use client';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { UserForm } from "@/components/user-form";
import {toast} from "react-toastify";
import { useRouter } from "next/navigation";

export default function NewUserPage() {
    const router = useRouter();

    const handleCreate = (data: any) => {
        // Here you would typically send a request to your backend to create the user
        console.log('Creating user:', data);
        toast.success('User created successfully');
        router.push('/users');
    }

    return (
        <div className="space-y-4">
            <h1 className="text-2xl font-semibold tracking-tight">New User</h1>
            <Card>
                <CardHeader>
                    <CardTitle>Create a New User</CardTitle>
                </CardHeader>
                <CardContent>
                    <UserForm onSubmit={handleCreate} submitLabel="Create User" />
                </CardContent>
            </Card>
        </div>
    )
}