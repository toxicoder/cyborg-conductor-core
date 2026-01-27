/**
 * Node.js adapter glue for cyborg execution.
 * This exposes an async function that can be called by the Go daemon via gRPC.
 */

import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import { promisify } from 'util';
import { exec } from 'child_process';
import { Readable } from 'stream';

// Promisify exec for easier async/await usage
const execAsync = promisify(exec);

// Define the interface for the run result
export interface RunResult {
  stdout: string;
  stderr: string;
  returncode: number;
}

/**
 * Run a script with the given arguments
 * @param script - Path to the script to execute
 * @param args - Arguments to pass to the script
 * @returns Promise resolving to the execution result
 */
export async function runScript(script: string, args: string[]): Promise<RunResult> {
  try {
    // Execute the script with timeout (3 seconds as per requirements)
    const result = await execAsync(`${script} ${args.join(' ')}`, {
      timeout: 3000 // 3 second timeout
    });
    
    return {
      stdout: result.stdout,
      stderr: result.stderr,
      returncode: 0
    };
  } catch (error: any) {
    if (error.killed) {
      return {
        stdout: '',
        stderr: 'Script timed out after 3 seconds',
        returncode: -1
      };
    }
    
    return {
      stdout: '',
      stderr: error.message || 'Unknown error occurred',
      returncode: error.code || -1
    };
  }
}

/**
 * gRPC service implementation
 */
export class CyborgService {
  /**
   * Run a script via gRPC
   * @param call - gRPC call object
   * @param callback - gRPC callback
   */
  async runScript(call: any, callback: any) {
    try {
      const { script, args } = call.request;
      
      const result = await runScript(script, args);
      
      callback(null, {
        stdout: result.stdout,
        stderr: result.stderr,
        returncode: result.returncode
      });
    } catch (error) {
      callback(error, null);
    }
  }
}

// Export the main function for use
export { runScript };

// Example usage (for testing)
if (require.main === module) {
  // This would be called from the Go daemon via gRPC
  console.log('Node.js adapter loaded');
}