/* tslint:disable */
/* eslint-disable */
/**
 * HashiCorp Vault API
 * HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.
 *
 * The version of the OpenAPI document: 1.20.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 *
 * @export
 * @interface UiConfigDeleteCustomMessageResponse
 */
export interface UiConfigDeleteCustomMessageResponse {
  /**
   *
   * @type {string}
   * @memberof UiConfigDeleteCustomMessageResponse
   */
  id?: string;
}

/**
 * Check if a given object implements the UiConfigDeleteCustomMessageResponse interface.
 */
export function instanceOfUiConfigDeleteCustomMessageResponse(
  value: object
): value is UiConfigDeleteCustomMessageResponse {
  return true;
}

export function UiConfigDeleteCustomMessageResponseFromJSON(json: any): UiConfigDeleteCustomMessageResponse {
  return UiConfigDeleteCustomMessageResponseFromJSONTyped(json, false);
}

export function UiConfigDeleteCustomMessageResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): UiConfigDeleteCustomMessageResponse {
  if (json == null) {
    return json;
  }
  return {
    id: json['id'] == null ? undefined : json['id'],
  };
}

export function UiConfigDeleteCustomMessageResponseToJSON(json: any): UiConfigDeleteCustomMessageResponse {
  return UiConfigDeleteCustomMessageResponseToJSONTyped(json, false);
}

export function UiConfigDeleteCustomMessageResponseToJSONTyped(
  value?: UiConfigDeleteCustomMessageResponse | null,
  ignoreDiscriminator: boolean = false
): any {
  if (value == null) {
    return value;
  }

  return {
    id: value['id'],
  };
}
