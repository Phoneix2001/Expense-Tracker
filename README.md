# Expense Tracker CLI

**Expense Tracker CLI** is a simple and fast command-line tool written in Go to help you keep track of your personal expenses from the terminal. Add, list, delete, and summarize your spending with just a few commands. All data is stored locally in an `expenses.json` file.

---

## 💰 Features

- ➕ Add new expense entries with description and amount  
- 📋 List all recorded expenses  
- 🗑️ Delete expenses by ID  
- 📈 Get total or monthly summaries  
- 📆 Automatically tracks date of each expense  
- ⚡ Fast and lightweight CLI built in Go  

---

## 📦 Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or later)

---

## 🔧 Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/expense-tracker.git
   cd expense-tracker
   ```

2. **Build the binary:**

   ```bash
   go build -o expense-tracker
   ```

3. *(Optional)* **Install globally:**

   ```bash
   sudo mv expense-tracker /usr/local/bin/
   ```

   You can now use `expense-tracker` from anywhere on your system.

---

## 🛠️ Usage

### ➕ Adding an expense

```bash
expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)
```

### 📋 Listing expenses

```bash
expense-tracker list
# ID  Date        Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10
```

### 🗑️ Deleting an expense

```bash
expense-tracker delete --id 2
# Expense deleted successfully
```

### 📈 Viewing summary

#### Total summary:

```bash
expense-tracker summary
# Total expenses: $20
```

#### Monthly summary:

```bash
expense-tracker summary --month 8
# Total expenses for August: $20
```

---

## 💡 Example Workflow

```bash
expense-tracker add --description "Lunch" --amount 20
expense-tracker add --description "Dinner" --amount 10
expense-tracker list
expense-tracker summary
expense-tracker delete --id 2
expense-tracker summary
expense-tracker summary --month 8
```

---

## 🌐 Project Source

This project follows the hands-on Go CLI project guidelines from [roadmap.sh](https://roadmap.sh/projects/expense-tracker).
