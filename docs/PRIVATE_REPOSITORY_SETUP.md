# Private Repository Setup for Sunday Schemas

> **Guide for using sunday-schemas Go module from private GitHub repository**

## 🔐 Issue: Private Repository Access

The error you're seeing occurs because:
1. The GitHub repository is private
2. Go module proxy cannot access private repositories
3. Git authentication is not configured for Go modules

## 🛠️ Solutions

### Option 1: Configure GOPRIVATE (Recommended)

Tell Go to bypass the module proxy for private repositories:

```bash
# Set GOPRIVATE to bypass proxy for this repository
go env -w GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas

# Alternative: Set for all repos under your account
go env -w GOPRIVATE=github.com/rakeyshgidwani/*

# Alternative: Add to existing GOPRIVATE setting
go env -w GOPRIVATE="$GOPRIVATE,github.com/rakeyshgidwani/sunday-schemas"
```

Then try the install again:
```bash
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

### Option 2: Configure Git Authentication

#### Using Personal Access Token (Recommended)
```bash
# Configure git to use token for GitHub
git config --global url."https://<USERNAME>:<TOKEN>@github.com/".insteadOf "https://github.com/"

# Replace <USERNAME> with your GitHub username
# Replace <TOKEN> with your personal access token
```

#### Using SSH (Alternative)
```bash
# Configure git to use SSH for GitHub
git config --global url."git@github.com:".insteadOf "https://github.com/"

# Make sure your SSH key is added to GitHub
ssh -T git@github.com
```

### Option 3: Environment Variables

Set environment variables for the current session:
```bash
export GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas
export GOSUMDB=off  # Disable checksum verification for private repos
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

## 🔍 Verification Steps

After configuring, verify the setup:

```bash
# Check GOPRIVATE setting
go env GOPRIVATE

# Test repository access
git ls-remote https://github.com/rakeyshgidwani/sunday-schemas.git

# Try installing the module
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1

# Test import in a simple program
```

## 🧪 Test Installation

Create a test program to verify the installation:

```bash
# Create test directory
mkdir test-sunday-schemas
cd test-sunday-schemas

# Initialize Go module
go mod init test-sunday-schemas

# Create test file
cat > main.go << 'EOF'
package main

import (
    "fmt"
    schemas "github.com/rakeyshgidwani/sunday-schemas/codegen/go"
)

func main() {
    fmt.Println("Sunday Schemas Go module loaded successfully!")
    fmt.Println("Available venues:", schemas.AllVenues())
    fmt.Println("Available schemas:", schemas.AllSchemas())
}
EOF

# Install dependency
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1

# Run test
go run main.go
```

Expected output:
```
Sunday Schemas Go module loaded successfully!
Available venues: [polymarket kalshi]
Available schemas: [raw.v0 md.orderbook.delta.v1 md.trade.v1 insights.arb.lite.v1 insights.movers.v1 insights.whales.lite.v1 insights.unusual.v1 infra.venue_health.v1]
```

## 🏢 For Team/Organization Setup

### Method 1: Team GOPRIVATE Configuration
Add to team development setup instructions:
```bash
# Add to ~/.bashrc or ~/.zshrc
export GOPRIVATE=github.com/rakeyshgidwani/*
```

### Method 2: Project-specific .envrc (with direnv)
```bash
# Create .envrc in project root
echo 'export GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas' > .envrc
direnv allow
```

### Method 3: Docker Development
```dockerfile
# In Dockerfile
ENV GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas

# Configure git credentials in container
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
```

## 🔧 Alternative: Make Repository Public

If the repository should be public, you can:

1. **Make the repository public** on GitHub:
   - Go to repository Settings
   - Scroll to "Danger Zone"
   - Click "Change repository visibility"
   - Select "Make public"

2. **Re-test the Go module**:
   ```bash
   # Clear Go module cache
   go clean -modcache

   # Install without GOPRIVATE
   go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
   ```

## 🚨 Troubleshooting

### Issue: Still getting 404 errors
**Solution**: Clear Go module cache and try again
```bash
go clean -modcache
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

### Issue: Permission denied
**Solution**: Check GitHub token permissions
```bash
# Test GitHub access
curl -H "Authorization: token YOUR_TOKEN" https://api.github.com/user
```

### Issue: SSH key issues
**Solution**: Add SSH key to GitHub
```bash
# Generate new SSH key if needed
ssh-keygen -t ed25519 -C "your_email@example.com"

# Add to ssh-agent
ssh-add ~/.ssh/id_ed25519

# Test SSH connection
ssh -T git@github.com
```

### Issue: Corporate firewall/proxy
**Solution**: Configure Go proxy settings
```bash
# Configure proxy if needed
go env -w GOPROXY=https://your-corporate-proxy.com,direct
go env -w GOPRIVATE=github.com/rakeyshgidwani/*
```

## 📋 Quick Fix Summary

**Fastest solution for most users:**

```bash
# Step 1: Set GOPRIVATE
go env -w GOPRIVATE=github.com/rakeyshgidwani/sunday-schemas

# Step 2: Configure git authentication (choose one)
# Option A: Personal Access Token
git config --global url."https://YOUR_USERNAME:YOUR_TOKEN@github.com/".insteadOf "https://github.com/"

# Option B: SSH
git config --global url."git@github.com:".insteadOf "https://github.com/"

# Step 3: Install module
go get github.com/rakeyshgidwani/sunday-schemas/codegen/go@v1.0.1
```

## 📞 Getting Help

If you continue having issues:

1. **Check your GitHub access**:
   ```bash
   git clone https://github.com/rakeyshgidwani/sunday-schemas.git
   ```

2. **Verify Go configuration**:
   ```bash
   go env GOPRIVATE
   go env GOPROXY
   ```

3. **Contact repository owner** for access permissions

4. **Use validation script** to test installation:
   ```bash
   ./scripts/validate-deployment.sh --version 1.0.1 --verbose
   ```

---

**Note**: These setup steps are one-time configuration. Once configured, the Go module will work normally for all team members.