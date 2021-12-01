/* tslint:disable */
/* eslint-disable */
/**
 * Mincuru Cars App
 * Mincuru Cars App API仕様
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: awwa500@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserResponse
 */
export interface UserResponse {
    /**
     * 
     * @type {number}
     * @memberof UserResponse
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof UserResponse
     */
    role: UserResponseRoleEnum;
    /**
     * 
     * @type {Date}
     * @memberof UserResponse
     */
    createdAt: Date;
    /**
     * 
     * @type {Date}
     * @memberof UserResponse
     */
    updatedAt: Date;
}

/**
* @export
* @enum {string}
*/
export enum UserResponseRoleEnum {
    User = 'user',
    Admin = 'admin'
}

export function UserResponseFromJSON(json: any): UserResponse {
    return UserResponseFromJSONTyped(json, false);
}

export function UserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'email': json['email'],
        'role': json['role'],
        'createdAt': (new Date(json['created_at'])),
        'updatedAt': (new Date(json['updated_at'])),
    };
}

export function UserResponseToJSON(value?: UserResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'email': value.email,
        'role': value.role,
        'created_at': (value.createdAt.toISOString()),
        'updated_at': (value.updatedAt.toISOString()),
    };
}

