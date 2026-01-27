#!/usr/bin/env python3
"""
Python adapter glue for cyborg execution.
This exposes a simple JSON RPC interface that can be called by the Go daemon.
"""

import json
import sys
import subprocess
import argparse
from typing import Dict, Any, Tuple

class CyborgExecutor:
    """Executes cyborg scripts via JSON RPC interface."""
    
    def __init__(self):
        pass
    
    def run_script(self, script: str, args: list) -> Dict[str, Any]:
        """
        Run a script with given arguments.
        
        Args:
            script: Path to the script to execute
            args: List of arguments to pass to the script
            
        Returns:
            Dictionary with result containing stdout, stderr, and exit code
        """
        try:
            # Execute the script with timeout (3 seconds as per requirements)
            result = subprocess.run(
                [script] + args,
                capture_output=True,
                text=True,
                timeout=3  # 3 second timeout
            )
            
            return {
                "stdout": result.stdout,
                "stderr": result.stderr,
                "returncode": result.returncode
            }
        except subprocess.TimeoutExpired:
            return {
                "stdout": "",
                "stderr": "Script timed out after 3 seconds",
                "returncode": -1
            }
        except Exception as e:
            return {
                "stdout": "",
                "stderr": str(e),
                "returncode": -1
            }
    
    def handle_request(self, request: Dict[str, Any]) -> Dict[str, Any]:
        """
        Handle a JSON RPC request.
        
        Args:
            request: JSON RPC request dictionary
            
        Returns:
            JSON RPC response dictionary
        """
        method = request.get("method")
        params = request.get("params", {})
        
        if method == "run":
            result = self.run_script(
                params.get("script", ""),
                params.get("args", [])
            )
            
            return {
                "id": request.get("id", 0),
                "jsonrpc": "2.0",
                "result": result
            }
        else:
            return {
                "id": request.get("id", 0),
                "jsonrpc": "2.0",
                "error": {
                    "code": -32601,
                    "message": "Method not found"
                }
            }

def main():
    """Main entry point for the Python adapter."""
    executor = CyborgExecutor()
    
    # Read request from stdin
    input_data = sys.stdin.read().strip()
    
    if not input_data:
        print(json.dumps({
            "error": {
                "code": -32700,
                "message": "Parse error"
            }
        }))
        return
    
    try:
        request = json.loads(input_data)
        response = executor.handle_request(request)
        print(json.dumps(response))
    except json.JSONDecodeError:
        print(json.dumps({
            "error": {
                "code": -32700,
                "message": "Parse error"
            }
        }))

if __name__ == "__main__":
    main()