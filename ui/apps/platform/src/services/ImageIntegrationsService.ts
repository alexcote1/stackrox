import axios from './instance';

import { IntegrationBase, IntegrationOptions } from './IntegrationsService';

const imageIntegrationsUrl = '/v1/imageintegrations';

const updateIntegrationKey = 'config';

// See proto/storage/image_integration.proto

export type ImageIntegrationCategory = 'REGISTRY' | 'SCANNER' | 'NODE_SCANNER';

export type ImageIntegrationBase = {
    clusters?: string[];
    categories: ImageIntegrationCategory[];
    autogenerated: boolean;
    clusterId: string;
    skipTestIntegration: boolean;
} & IntegrationBase;

/*
 * Create integration.
 * The id of argument is empty string and the id of response is assigned by server.
 */
export function createImageIntegration(
    integration: ImageIntegrationBase
): Promise<ImageIntegrationBase> {
    return axios.post(imageIntegrationsUrl, integration);
}

/*
 * Read integrations (plural).
 */
export function fetchImageIntegrations(): Promise<ImageIntegrationBase[]> {
    return axios
        .get<{ integrations: ImageIntegrationBase[] }>(imageIntegrationsUrl)
        .then((response) => response.data.integrations);
}

/*
 * Update integration.
 *
 * Call with options argument if integration has stored credentials, aka password:
 * true to update credentials on the server from the request payload
 * false not to update credentials on the server
 *
 * Call without options argument if integration does not have stored credentials.
 */
export function saveImageIntegration(
    integration: ImageIntegrationBase,
    { updatePassword }: IntegrationOptions = {}
): Promise<Record<string, never>> {
    const { id } = integration;

    if (!id) {
        throw new Error('Integration entity must have an id to be saved');
    }

    if (typeof updatePassword === 'boolean') {
        return axios.patch(`${imageIntegrationsUrl}/${id}`, {
            [updateIntegrationKey]: integration,
            updatePassword,
        });
    }

    return axios.put(`${imageIntegrationsUrl}/${id}`, integration);
}

/*
 * Test integration.
 *
 * Call with options argument if integration has stored credentials, aka password:
 * true to use credentials in the request payload
 * false to use credentials on the server
 *
 * Call without options argument if integration does not have stored credentials.
 */
export function testImageIntegration(
    integration: ImageIntegrationBase,
    { updatePassword }: IntegrationOptions = {}
): Promise<Record<string, never>> {
    if (typeof updatePassword === 'boolean') {
        return axios.post(`${imageIntegrationsUrl}/test/updated`, {
            [updateIntegrationKey]: integration,
            updatePassword,
        });
    }

    return axios.post(`${imageIntegrationsUrl}/test`, integration);
}

/*
 * Delete integration (singular).
 */
export function deleteImageIntegration(id: string): Promise<Record<string, never>> {
    return axios.delete(`${imageIntegrationsUrl}/${id}`);
}

/*
 * Delete integrations (plural).
 */
export function deleteImageIntegrations(ids: string[]): Promise<Record<string, never>[]> {
    return Promise.all(ids.map((id) => deleteImageIntegration(id)));
}
