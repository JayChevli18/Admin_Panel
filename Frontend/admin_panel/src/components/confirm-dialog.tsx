'use client';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

export const ConfirmDialog = ({
    isOpen,
    title='Confirm',
    description='Are you sure you want to proceed?',
    confirmText='Delete',
    cancelText='Cancel',
    onConfirm,
    onCancel
}: {
    isOpen: boolean;
    title: string;
    description: string;
    confirmText?: string;
    cancelText?: string;
    onConfirm: () => void;
    onCancel: () => void;
}) => {
    return (
        <Dialog open={isOpen} onOpenChange={(v)=>!v && onCancel()}>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>{title}</DialogTitle>
                    <DialogDescription>{description}</DialogDescription>
                </DialogHeader>
                <DialogFooter>
                    <Button variant="ghost" onClick={onCancel}>{cancelText}</Button>
                    <Button variant="destructive" onClick={onConfirm}>{confirmText}</Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}