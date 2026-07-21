import { Suspense } from "react";
import WebhookForm from "./component/webhookform"; 

export default function SetupPage() {
    return (
        <main className="min-h-screen flex items-center justify-center">
            {}
            <Suspense fallback={<p>Cargando configuración...</p>}>
                <WebhookForm />
            </Suspense>
        </main>
    );
}